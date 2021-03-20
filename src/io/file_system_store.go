package main

import (
	"io"
)

type FileSystemPlayerStore struct {
	database io.ReadSeeker
}

// GetLeague loads the JSON data from database and returns it
func (f *FileSystemPlayerStore) GetLeague() []Player {
	_, _ = f.database.Seek(0, io.SeekStart)
	league, _ := LoadLeague(f.database)
	return league
}

// FileSystemPlayerStore checks if name contains any entries
// and returns (Wins, true) if so, else (0, false)
func (f *FileSystemPlayerStore) GetPlayerScore(name string) (int, bool) {
	var wins int
	found := false

	for _, player := range f.GetLeague() {
		if player.Name == name {
			found = true
			wins = player.Wins
			break
		}
	}

	return wins, found
}
