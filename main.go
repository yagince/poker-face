package main

import (
	"./poker"
	"fmt"
	"net/http"
)

func main() {
	// websocket
	server := poker.NewServer("/entry")
	go server.Listen()

	http.Handle("/", http.FileServer(http.Dir(".")))
	http.ListenAndServe(":8080", nil)
}
