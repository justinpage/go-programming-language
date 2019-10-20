package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	const guide = `<p>A: {{.A}}</p><p>B: {{.B}}</p>`

	var report = template.Must(template.New("escape").Parse(guide))

	var data struct {
		A string        // untrusted plain text
		B template.HTML // trusted HTML
	}

	data.A = "<b>Hello!</b>"
	data.B = "<b>Hello!</b>"

	// run: go run autoescape  > autoescape.html
	if err := report.Execute(os.Stdout, data); err != nil {
		log.Fatal(err)
	}
}
