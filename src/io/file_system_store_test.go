package main

import (
	"strings"
	"testing"
)

func TestFileSystemStore(t *testing.T) {
	database := strings.NewReader(`[
            {"Name": "Cleo", "Wins": 10},
            {"Name": "Chris", "Wins": 33}]`)
	store := FileSystemPlayerStore{database}

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

		assertStoreEquals(t, got, 33)

	})
}

func assertStoreEquals(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Fatalf("got %d want %d", got, want)
	}
}
