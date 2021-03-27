package main

import (
	"encoding/json"
	"fmt"
	"io"
)

type League []Player

// Find finds a Player record given its name
// and returns a pointer to it if found, else nil.
func (l League) Find(name string) *Player {
	for i, p := range l {
		if p.Name == name {
			return &l[i]
		}
	}
	return nil
}

// LoadLegaue loads a league from JSON
func LoadLeague(rdr io.Reader) (League, error) {
	var league League
	err := json.NewDecoder(rdr).Decode(&league)
	if err != nil {
		err = fmt.Errorf("problem parsing league, %v", err)
	}
	return league, err
}