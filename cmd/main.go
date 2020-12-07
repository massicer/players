package main

import (
	"net/http"
	"log"
	"github.com/massicer/players/internal/server"
	"github.com/massicer/players/internal/store"
)

func main() {
    server := &server.PlayerServer{&store.InMemoryPlayerStore{}}

    if err := http.ListenAndServe(":5000", server); err != nil {
        log.Fatalf("could not listen on port 5000 %v", err)
    }
}