package main

import (
	"bytes"
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type WebSocketServer struct {
	http.Handler
}

func (s WebSocketServer) homePage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "client.html")
}

func (s WebSocketServer) webSocket(w http.ResponseWriter, r *http.Request) {
	ws, _ := upgrader.Upgrade(w, r, nil)

	for {
		messageType, message, _ := ws.ReadMessage()
		if messageType != websocket.TextMessage {
			ws.Close()

			return
		}

		message = bytes.Replace(message, []byte("?"), []byte("!"), -1)
		ws.WriteMessage(websocket.TextMessage, message)
	}
}

func NewServer() *WebSocketServer {
	server := new(WebSocketServer)

	router := http.NewServeMux()
	router.Handle("/", http.HandlerFunc(server.homePage))
	router.Handle("/ws", http.HandlerFunc(server.webSocket))

	server.Handler = router

	return server
}
