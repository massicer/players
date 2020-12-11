package server

import (
	"net/http"
    "fmt"
    "github.com/massicer/players/internal/entities"
    "encoding/json"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
    RecordWin(name string) error
    GetLeagueTable() []entities.Player
}


type PlayerServer struct {
    Store PlayerStore
    http.Handler
}

func NewPlayerServer(store PlayerStore) *PlayerServer {
    p := new(PlayerServer)

    p.Store = store

    router := http.NewServeMux()
    router.Handle("/league", http.HandlerFunc(p.leagueHandler))
    router.Handle("/players/", http.HandlerFunc(p.playersHandler))

    p.Handler = router

    return p
}

func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
    leagueTable := p.Store.GetLeagueTable()
    err := json.NewEncoder(w).Encode(leagueTable)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
    }
    w.WriteHeader(http.StatusOK)
}

func (p *PlayerServer) playersHandler(w http.ResponseWriter, r *http.Request) {
    player := r.URL.Path[len("/players/"):]

    switch r.Method {
    case http.MethodPost:
        p.processWin(w, player)
    case http.MethodGet:
        p.showScore(w, player)
    }
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
    }
	w.WriteHeader(http.StatusAccepted)
}