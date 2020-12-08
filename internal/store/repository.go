package store

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