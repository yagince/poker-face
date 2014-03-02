package main

import (
	"net/http"
	"fmt"
	"./poker"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "path : %s", r.URL.Path)
}

func main() {
	// test
	http.HandleFunc("/hello", helloHandler)

	// websocket
	server := poker.NewServer("/entry")
	go server.Listen()

	http.Handle("/", http.FileServer(http.Dir(".")))
	// for _, path := range []string{"lib", "js"} {
	// 	http.Handle("/"+path+"/", http.StripPrefix("/"+path+"/", http.FileServer(http.Dir(path))))
	// }
	http.ListenAndServe(":8080", nil)
}
