package main

import (
	"flag"
	"log"
	"strings"

	"github.com/gen2brain/beeep"
)

func main() {
	var title string
	flag.StringVar(&title, "title", "tinynotify", "title of notification")
	flag.Parse()
	message := strings.Join(flag.Args(), "\n\n")

	if err := run(title, message); err != nil {
		log.Fatalf("!! %+v", err)
	}
}

// TODO: use embed (assets)

func run(title, message string) error {
	return beeep.Notify(title, message, "assets/information.png")
}
