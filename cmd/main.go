package main

import (
	"net/http"
	"log"
	"fmt"
	"github.com/massicer/players/internal/server"
	"github.com/massicer/players/internal/store"
)

const server_port int = 5000

func main() {
	log.Printf("Preparing to start server on port: %d", server_port)
    server := &server.PlayerServer{
		Store: &store.InMemoryPlayerStore{
			Scores: make(map[string]int),
		},
	}

	if err := http.ListenAndServe(fmt.Sprintf(":%d", server_port), server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
    
}