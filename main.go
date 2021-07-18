package main

import (
	"log"
	"os"

	"github.com/gen2brain/beeep"
)

func main() {
	title := os.Args[1]
	message := os.Args[2]
	if err := run(title, message); err != nil {
		log.Fatalf("!! %+v", err)
	}
}

// TODO: use embed (assets)

func run(title, message string) error {
	return beeep.Notify(title, message, "assets/information.png")
}
