package server

import (
	"net/http"
	"fmt"
	
)

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string) error
}

type PlayerServer struct {
    Store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

    router := http.NewServeMux()

    router.Handle("/league", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
    }))

    router.Handle("/players/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        player := r.URL.Path[len("/players/"):]

        switch r.Method {
        case http.MethodPost:
            p.processWin(w, player)
        case http.MethodGet:
            p.showScore(w, player)
        }
    }))

    router.ServeHTTP(w, r)
}

func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {
    score := p.Store.GetPlayerScore(player)
  
    if score == 0 {
        w.WriteHeader(http.StatusNotFound)
    }

    fmt.Fprint(w, score)
}

func (p *PlayerServer) processWin(w http.ResponseWriter, player string) {
    err := p.Store.RecordWin(player)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprint(w, "Cannot process win")
    } else {
        w.WriteHeader(http.StatusAccepted)
    }
}