package main

import (
	"log"
	"os"
	"text/template"
	"time"

	"github.com/justinpage/go-programming-language/4-composite-types/examples/10.json/github"
)

func main() {
	const issueList = `{{.TotalCount}} issues
	{{range .Items}}----------------------------------------
	Number: {{.Number}}
	User: {{.User.Login}}
	Title: {{.Title | printf "%.64s"}}
	Age: {{.CreatedAt | daysAgo}} days
	{{end}}` + "\n"

	var report = template.Must(template.New("issuelist").
		Funcs(template.FuncMap{"daysAgo": daysAgo}).
		Parse(issueList))

	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}
