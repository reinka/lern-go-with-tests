package main

import (
	"encoding/json"
	"fmt"
	"io"
)

// LoadLegaue loads a league from JSON
func LoadLeague(rdr io.Reader) ([]Player, error) {
	var league []Player
	err := json.NewDecoder(rdr).Decode(&league)
	if err != nil {
		err = fmt.Errorf("problem parsing league, %v", err)
	}
	return league, err
}
