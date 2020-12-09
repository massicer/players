package store

import "github.com/massicer/players/internal/entities"

type InMemoryPlayerStore struct{
    Scores map[string]int
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
    return i.Scores[name]
}

func (i *InMemoryPlayerStore) RecordWin(name string) error{
    i.Scores[name]++
    return nil
}

func (i *InMemoryPlayerStore) GetLeagueTable() []entities.Player{
    return []entities.Player {
        entities.Player{Name: "Max", Wins: 0},
    }
}