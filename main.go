package main

import (
	"./poker"
	"flag"
	"fmt"
	"net/http"
)

const (
	DefaultPort = 8080
)

func main() {
	// websocket
	server := poker.NewServer("/entry")
	go server.Listen()

	http.Handle("/", http.FileServer(http.Dir(".")))

	var port int
	flag.IntVar(&port, "port", DefaultPort, "port number")
	flag.IntVar(&port, "p", DefaultPort, "port number")
	flag.Parse()

	fmt.Printf("listen port:%d\n", port)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		fmt.Println(err)
	}
}
