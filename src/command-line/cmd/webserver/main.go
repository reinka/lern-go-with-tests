package main

import (
	"github.com/reinka/lern-go-with-tests/src/command-line"
	"log"
	"net/http"
)

const dbFileName = "game.db.json"

func main() {
	store, closeFn, err := poker.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}
	defer closeFn()

	server := poker.NewPlayerServer(store)

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
