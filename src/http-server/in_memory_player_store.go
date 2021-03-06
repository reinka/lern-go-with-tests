package main

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}

type InMemoryPlayerStore struct {
	store map[string]int
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) (int, bool) {
	score, found := i.store[name]
	return score, found
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.store[name]++
}
