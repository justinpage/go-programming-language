package main

import (
	"fmt"
	"os"
	"reflect"

	"github.com/justinpage/go-programming-language/12-reflection/examples/2.display/display"
	"github.com/justinpage/go-programming-language/12-reflection/examples/2.display/eval"
)

func main() {
	e, _ := eval.Parse("sqrt(A / pi)")
	display.Display("e", e)
	fmt.Println("")

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

	display.Display("strangelove", strangelove)
	fmt.Println("")

	display.Display("os.Stderr", os.Stderr)
	fmt.Println("")

	display.Display("rV", reflect.ValueOf(os.Stderr))
}
