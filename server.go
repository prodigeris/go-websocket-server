package main

import (
	"fmt"
	"net/http"
)

type WebSocketServer struct {
	http.Handler
}

func (s WebSocketServer) homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hi!")
}

func NewServer() *WebSocketServer {
	server := new(WebSocketServer)

	router := http.NewServeMux()
	router.Handle("/", http.HandlerFunc(server.homePage))

	server.Handler = router

	return server
}
