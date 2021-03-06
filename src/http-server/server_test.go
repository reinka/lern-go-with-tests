package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
)

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
}

func (s *StubPlayerStore) GetPlayerScore(name string) (int, bool) {
	score, found := s.scores[name]
	return score, found
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
	s.scores[name]++
}

func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"Floyd":  10,
			"Pepper": 20,
			"Salt":   0,
		},
		nil,
	}
	server := &PlayerServer{store: &store}

	t.Run("get Pepper's score", func(t *testing.T) {
		request := newGetScoreRequest("Pepper")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "20")
	})

	t.Run("get Floyd's score", func(t *testing.T) {
		request := newGetScoreRequest("Floyd")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "10")
	})

	t.Run("get 404 on missing players", func(t *testing.T) {
		request := newGetScoreRequest("Apollo")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusNotFound)
	})

	t.Run("get player score with value 0", func(t *testing.T) {
		request := newGetScoreRequest("Salt")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
	})
}

func TestStoreWins(t *testing.T) {

	t.Run("record wins when POST", func(t *testing.T) {
		store := StubPlayerStore{
			map[string]int{},
			nil,
		}
		server := &PlayerServer{store: &store}

		player := "Pepper"
		request := newPostWinRequest(player)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusAccepted)
		if len(store.winCalls) != 1 {
			t.Errorf("got %d calls to RecordWin, want %d", len(store.winCalls), 1)
		}

		if store.winCalls[0] != player {
			t.Errorf("did not store the correct winner, got %q, want %q", store.winCalls[0], player)
		}
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		store := StubPlayerStore{
			map[string]int{},
			nil,
		}
		server := &PlayerServer{store: &store}

		player := "Pepper"
		wantedWins := 1000

		var wg sync.WaitGroup
		wg.Add(wantedWins)

		// do all the POST request concurrently
		for i := 0; i < wantedWins; i++ {
			go func(w *sync.WaitGroup) {
				server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
				w.Done()
			}(&wg)
		}
		wg.Wait()

		// GET the players score
		response := httptest.NewRecorder()

		server.ServeHTTP(response, newGetScoreRequest(player))

		assertStatus(t, response.Code, http.StatusOK)

		if len(store.winCalls) != wantedWins {
			t.Errorf("got %d winCalls, want %d", len(store.winCalls), wantedWins)
		}
	})
}

func newGetScoreRequest(name string) *http.Request {
	return httptest.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
}

func newPostWinRequest(name string) *http.Request {
	return httptest.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got status %d, want %d", got, want)
	}
}
