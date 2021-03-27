package main

import (
	"testing"
)

func TestFileSystemStore(t *testing.T) {
	database, removeTmpFile := createTempFile(t, `[
            {"Name": "Cleo", "Wins": 10},
            {"Name": "Chris", "Wins": 33}]`)
	defer removeTmpFile()
	store, err := NewFileSystemPlayerStore(database)

	assertNoError(t, err)

	t.Run("league from a reader", func(t *testing.T) {

		got := store.GetLeague()

		want := []Player{
			{"Cleo", 10},
			{"Chris", 33},
		}

		assertLeague(t, got, want)

		// read again
		got = store.GetLeague()
		assertLeague(t, got, want)
	})

	t.Run("get score", func(t *testing.T) {
		got, _ := store.GetPlayerScore("Chris")
		assertScoreEquals(t, got, 33)
	})

	t.Run("store win", func(t *testing.T) {
		store.RecordWin("Chris")
		got, _ := store.GetPlayerScore("Chris")
		assertScoreEquals(t, got, 34)
	})

	t.Run("store wins for new player", func(t *testing.T) {
		store.RecordWin("Shibe")
		got, _ := store.GetPlayerScore("Shibe")
		assertScoreEquals(t, got, 1)
	})
}

func assertScoreEquals(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Fatalf("got %d want %d", got, want)
	}
}

//assertNoError asserts that no error occurred
func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("didn't expect an error but got one, %v", err)
	}
}
