package poker

import (
	"code.google.com/p/go.net/websocket"
	"net/http"
	"log"
)

// Chat server.
type Server struct {
	path string
	clients []*Client
	addClient chan *Client
	removeClient chan *Client
	sendAll chan *Choice
	choices []*Choice
}

// Create new chat server.
func NewServer(path string) *Server {
	clients := make([]*Client, 0)
	addClient := make(chan *Client)
	removeClient := make(chan *Client)
	sendAll := make(chan *Choice)
	choices := make([]*Choice, 0)
	return &Server{path, clients, addClient, removeClient, sendAll, choices}
}

func (self *Server) AddClient() chan<- *Client {
	return (chan<- *Client)(self.addClient)
}

func (self *Server) RemoveClient() chan<- *Client {
	return (chan<- *Client)(self.removeClient)
}

func (self *Server) SendAll() chan<-*Choice {
	return (chan<-*Choice)(self.sendAll)
}

func (self *Server) Choices() []*Choice {
	msgs := make([]*Choice, len(self.choices))
	copy(msgs, self.choices)
	return msgs
}

func (self *Server) Reset() {
	self.choices = make([]*Choice, 0)
}

// Listen and serve.
// It serves client connection and broadcast request.
func (self *Server) Listen() {

	log.Println("Listening server...")

	// websocket handler
	onConnected := func(ws *websocket.Conn) {
		client := NewClient(ws, self)
		self.addClient <- client
		client.Listen()
		defer ws.Close()
	}
	http.Handle(self.path, websocket.Handler(onConnected))
	log.Println("Created handler")

	for {
		select {

		// Add new a client
		case c := <-self.addClient:
			log.Println("Added new client")
			self.clients = append(self.clients, c)
			for _, msg := range self.choices {
				c.Write() <- msg
			}
			log.Println("Now", len(self.clients), "clients connected.")

		// remove a client
		case c := <-self.removeClient:
			log.Println("Remove client")
			for i := range self.clients {
				if self.clients[i] == c {
					self.clients = append(self.clients[:i], self.clients[i+1:]...)
					break
				}
			}

		// broadcast message for all clients
		case msg := <-self.sendAll:
			log.Println("Send all:", msg)
			if msg.ShouldSave() { self.choices = append(self.choices, msg) }
			for _, c := range self.clients {
				c.Write() <- msg
			}
		}
	}
}
