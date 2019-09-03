package main

import (
	"fmt"
	"net/http"
)

func WebSocketServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hi!")
}
