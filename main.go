package main

import (
	"context"
	"log"
	"net/http"
)

func main() {
	setUpApi()
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func setUpApi() {
	ctx := context.Background()

	manager := NewManager(ctx)

	http.Handle("/", http.FileServer(http.Dir("./frontent")))
	http.HandleFunc("/ws", manager.serveWS)
	http.HandleFunc("/login",manager.loginHandler)
}
