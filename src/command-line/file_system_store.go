package poker

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"testing"
)

type FileSystemPlayerStore struct {
	database *json.Encoder
	league   League
}

// NewFileSystemPlayerStore creates a new FileSystemPlayerStore and returns a
// pointer to it and an error, if any
func NewFileSystemPlayerStore(file *os.File) (*FileSystemPlayerStore, error) {
	err := initialisePlayerDBFile(file)

	if err != nil {
		return nil, fmt.Errorf("problem initialising player db file, %v", err)
	}

	league, err := NewLeague(file)

	if err != nil {
		return nil, fmt.Errorf(
			"problem loading player store from file %s, %v",
			file.Name(), err)
	}

	return &FileSystemPlayerStore{
		database: json.NewEncoder(&tape{file}),
		league:   league,
	}, err
}

func FileSystemPlayerStoreFromFile(path string) (*FileSystemPlayerStore, func(), error) {
	db, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		return nil, nil, fmt.Errorf("problem opening %s %v", path, err)
	}

	closeFunc := func() {
		_ = db.Close()
	}

	store, err := NewFileSystemPlayerStore(db)

	if err != nil {
		return nil, nil, fmt.Errorf("problem creating file system player store, %v ", err)
	}

	return store, closeFunc, nil
}

// initialisePlayerDBFile initialises a Player DB file
// and returns an error, if any
func initialisePlayerDBFile(file *os.File) error {
	_, err := file.Seek(0, 0)

	if err != nil {
		return fmt.Errorf("problem seeking file %s, %v",
			file.Name(), err)
	}

	info, err := file.Stat()

	if err != nil {
		return fmt.Errorf("problem getting file info from file %s, %v",
			file.Name(), err)
	}

	if info.Size() == 0 {
		_, _ = file.Write([]byte("[]"))
		_, _ = file.Seek(0, 0)
	}

	return nil
}

// GetLeague loads the JSON data from database, sorts its
// entries by their scores, from highest to lowest, and returns it
func (f *FileSystemPlayerStore) GetLeague() League {
	sort.Slice(f.league, func(i, j int) bool {
		return f.league[i].Wins > f.league[j].Wins
	})

	return f.league
}

// FileSystemPlayerStore checks if name contains any entries
// and returns (Wins, true) if so, else (0, false)
func (f *FileSystemPlayerStore) GetPlayerScore(name string) (int, bool) {
	player := f.league.Find(name)

	if player != nil {
		return player.Wins, true
	}

	return 0, false
}

// RecordWin increaments the Wins counter of a Player if it exists
// else it adds a new Player with Wins == 1. Additionally, it updates
// the league of the Player
func (f *FileSystemPlayerStore) RecordWin(name string) {
	player := f.league.Find(name)
	if player != nil {
		player.Wins++
	} else {
		f.league = append(f.league, Player{name, 1})
	}

	_ = f.database.Encode(f.league)
}

// createTempFile creates a temporary file under the default TempFile
// directory and returns it, together with a clean-up function
func createTempFile(t testing.TB, initialData string) (*os.File, func()) {
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
