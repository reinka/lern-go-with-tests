package main

import (
	"fmt"
	poker "github.com/reinka/lern-go-with-tests/src/command-line"
	"log"
	"os"
)

const dbFileName = "game.db.json"

func main() {
	fmt.Println("Let's play poker")
	fmt.Println("Type {Name} wins to record a win")

	store, closeFn, err := poker.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}
	defer closeFn()

	game := poker.NewCLI(store, os.Stdin)
	game.PlayPoker()
}
