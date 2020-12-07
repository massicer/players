package store

type InMemoryPlayerStore struct{}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
    return 123
}

func (i *InMemoryPlayerStore) RecordWin(name string) error{
    return nil
}