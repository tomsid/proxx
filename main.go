package main

import (
	"log"
	"os"

	"github.com/tomsid/proxx/proxx"
)

func main() {
	game, err := proxx.NewGame(5, 5, 5)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	game.Start()
}
