package main

import (
	"fmt"

	"github.com/justinpage/go-programming-language/12-reflection/exercises/12.1.display/display"
)

func main() {
	// Now that our Display function is complete, let's put it to work
	type Movie struct {
		Title, Subtitle string
		Year            int
		Color           bool
		Actors          map[string]string
		Oscars          []string
		Sequel          *string
	}

	// Let's declare a value of this type to see what Display does with it
	strangelove := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year:     1964,
		Color:    false,
		Actors: map[string]string{
			"Dr. Strangelove":            "Peter Sellers",
			"Grp. Capt. Lionel Mandrake": "Peter Sellers",
			"Pres. Merkin Muffley":       "Peter Sellers",
			"Gen. Buck Turgidson":        "George C. Scott",
			"Brig. Gen. Jack D. Ripper":  "Sterling Hayden",
			`Maj. T.J. "King" Kong`:      "Slim Pickens",
		},
		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screenplay (Nomin.)",
			"Best Director (Nomin.)",
			"Best Picture (Nomin.)",
		},
	}

	strangeloveTwo := Movie{
		Title:    "Dr. Strangelove 2",
		Subtitle: "How I Learned to Start Worrying and Hate the Bomb",
		Year:     1984,
		Color:    true,
		Actors: map[string]string{
			"Dr. Strangelove":            "Seter Pellers",
			"Grp. Capt. Lionel Mandrake": "Seter Pellers",
			"Pres. Merkin Muffley":       "Seter Pellers",
			"Gen. Buck Turgidson":        "Seorge C. Gcott",
			"Brig. Gen. Jack D. Ripper":  "Hterling Sayden",
			`Maj. T.J. "King" Kong`:      "Plim Sickens",
		},
		Oscars: []string{
			"Worst Actor (Nomin.)",
			"Worst Adapted Screenplay (Nomin.)",
			"Worst Director (Nomin.)",
			"Worst Picture (Nomin.)",
		},
	}

	movies := make(map[*Movie]bool)
	movies[&strangelove] = true
	movies[&strangeloveTwo] = false

	films := [2]Movie{strangelove, strangeloveTwo}

	display.Display("strangelove", strangelove)
	fmt.Println("")
	display.Display("movies", movies)
	fmt.Println("")
	display.Display("films", films)
}
