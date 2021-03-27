package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

type FileSystemPlayerStore struct {
	database io.ReadWriteSeeker
}

// GetLeague loads the JSON data from database and returns it
func (f *FileSystemPlayerStore) GetLeague() League {
	_, _ = f.database.Seek(0, io.SeekStart)
	league, _ := LoadLeague(f.database)
	return league
}

// FileSystemPlayerStore checks if name contains any entries
// and returns (Wins, true) if so, else (0, false)
func (f *FileSystemPlayerStore) GetPlayerScore(name string) (int, bool) {
	player := f.GetLeague().Find(name)

	if player != nil {
		return player.Wins, true
	}

	return 0, false
}

// RecordWin increaments the Wins counter of a Player if it exists
// else it adds a new Player with Wins == 1. Additionally, it updates
// the league of the Player
func (f *FileSystemPlayerStore) RecordWin(name string) {
	league := f.GetLeague()
	player := league.Find(name)
	if player != nil {
		player.Wins++
	} else {
		league = append(league, Player{name, 1})
	}

	_, _ = f.database.Seek(0, 0)
	_ = json.NewEncoder(f.database).Encode(league)
}

// createTempFile creates a temporary file under the default TempFile
// directory and returns it, together with a clean-up function
func createTempFile(t testing.TB, initialData string) (io.ReadWriteSeeker, func()) {
	t.Helper()

	tmpfile, err := ioutil.TempFile("", "db")
	if err != nil {
		t.Fatalf("could not create temp file %v", err)
	}

	_, _ = tmpfile.Write([]byte(initialData))

	removeFile := func() {
		_ = tmpfile.Close()
		_ = os.Remove(tmpfile.Name())
	}

	return tmpfile, removeFile
}
