package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"
)

type PlayerStore interface {
	GetPlayerScore(name string) (int, bool)
	RecordWin(name string)
	GetLeague() []Player
}

type Player struct {
	Name string
	Wins int
}

type PlayerServer struct {
	store  PlayerStore
	router *http.ServeMux
	mu     sync.Mutex
}

func NewPlayerServer(store PlayerStore) *PlayerServer {
	p := &PlayerServer{
		store:  store,
		router: http.NewServeMux(),
	}

	p.rountes()

	return p
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p.router.ServeHTTP(w, r)
}

func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {
	score, found := p.store.GetPlayerScore(player)

	if !found {
		w.WriteHeader(http.StatusNotFound)
	}

	_, _ = fmt.Fprint(w, score)
}

func (p *PlayerServer) handleLeague() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewEncoder(w).Encode(p.store.GetLeague())
		w.WriteHeader(http.StatusOK)
	}
}

func (p *PlayerServer) handlePlayers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		player := strings.TrimPrefix(r.URL.Path, "/players/")
		switch r.Method {
		case http.MethodPost:
			p.processWin(w, player)
		case http.MethodGet:
			p.showScore(w, player)
		}
	}
}

func (p *PlayerServer) processWin(w http.ResponseWriter, player string) {
	p.mu.Lock()
	p.store.RecordWin(player)
	p.mu.Unlock()

	w.WriteHeader(http.StatusAccepted)
}
