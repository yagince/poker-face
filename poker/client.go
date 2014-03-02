package poker

import (
	"code.google.com/p/go.net/websocket"
	"log"
)

// Chat client.
type Client struct {
	ws *websocket.Conn
	server *Server
	ch chan *Choice
	done chan bool
}

// write channel buffer size
const channelBufSize = 10

// Create new chat client.
func NewClient(ws *websocket.Conn, server *Server) *Client {

	if ws == nil {
		panic("ws cannot be nil")
	} else if server == nil {
		panic("server cannot be nil")
	}

	ch := make(chan *Choice, channelBufSize)
	done := make(chan bool)

	return &Client{ws, server, ch, done}
}

// Get websocket connection.
func (self *Client) Conn() *websocket.Conn {
	return self.ws
}

// Get Write channel
func (self *Client) Write() chan<-*Choice {
	return (chan<-*Choice)(self.ch)
}

// Get done channel.
func (self *Client) Done() chan<-bool {
	return (chan<-bool)(self.done)
}

// Listen Write and Read request via chanel
func (self *Client) Listen() {
	go self.listenWrite()
	self.listenRead()
}

// Listen write request via chanel
func (self *Client) listenWrite() {
	log.Println("Listening write to client")
	for {
		select {

		// send message to the client
		case choice := <-self.ch:
			log.Println("Send:", choice)
			websocket.JSON.Send(self.ws, choice)

		// receive done request
		case <-self.done:
			self.server.RemoveClient() <- self
			self.done <- true // for listenRead method
			return
		}
	}
}

// Listen read request via chanel
func (self *Client) listenRead() {
	log.Println("Listening read from client")
	for {
		select {

		// receive done request
		case <-self.done:
			self.server.RemoveClient() <- self
			self.done <- true // for listenWrite method
			return

		// read data from websocket connection
		default:
			log.Println("wait for receive message...")
			var choice Choice
			err := websocket.JSON.Receive(self.ws, &choice)
			if err != nil {
				self.done<-true
			} else {
				if choice.Reset { self.server.Reset() }
				self.server.SendAll() <- &choice
			}
		}
	}
}
