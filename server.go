package main

import (
	"fmt"
	"net/http"
)

type WebSocketServer struct {
}

func (s WebSocketServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hi!")
}

func NewServer() WebSocketServer {
	return WebSocketServer{}
}
