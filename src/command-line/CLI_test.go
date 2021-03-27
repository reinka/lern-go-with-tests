package poker_test

import (
	poker "github.com/reinka/lern-go-with-tests/src/command-line"
	"strings"
	"testing"
)

func TestCLI(t *testing.T) {

	t.Run("record chris win from user input", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")
		playerStore := newPlayerStore()

		cli := poker.NewCLI(playerStore, in)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, "Chris")
	})

	t.Run("record cleo win from user input", func(t *testing.T) {
		in := strings.NewReader("Cleo wins\n")
		playerStore := newPlayerStore()

		cli := poker.NewCLI(playerStore, in)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, "Cleo")
	})
}

func newPlayerStore() *poker.StubPlayerStore {
	return &poker.StubPlayerStore{
		map[string]int{
			"Chris": 10,
			"Cleo":  5,
		}, nil, nil}
}
