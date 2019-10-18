package main

import (
	"fmt"
	"log"
	"os"

	"github.com/justinpage/go-programming-language/4-composite-types/exercises/4.12.xkcd/comic"
)

func main() {
	command := os.Args[1:]

	if len(command) == 0 {
		log.Fatal("no command")
	}

	if command[0] == "index" {
		if err := xkcd.SeedIndexAlt(); err != nil {
			log.Fatal(err)
		}

		log.Print("index seeded")
		os.Exit(0)
	}

	comics, err := xkcd.Search(command[0])
	if err != nil {
		log.Fatal(err)
	}

	for _, comic := range comics {
		fmt.Printf("TITLE: %s\n", comic.Title)
		fmt.Printf("URL: %s\n", comic.URL)
		fmt.Printf("TRANSCRIPT:\n%s\n", comic.Transcript)
		fmt.Println("----------------------------------")
	}
}
