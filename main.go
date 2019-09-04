package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := flag.Uint("port", 1234, "Port of server")
	flag.Parse()

	if err := http.ListenAndServe(fmt.Sprintf(":%d", *port), newServer()); err != nil {
		log.Fatalf("Cannot start a server with port %d. %v", *port, err)
	}
}
