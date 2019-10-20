package main

import (
	"html/template"
	"log"
	"os"
	"time"

	"github.com/justinpage/go-programming-language/4-composite-types/examples/10.json/github"
)

func main() {
	const issueList = `
	<h1>{{.TotalCount}} issues</h1>
	<table>
	<tr style='text-align: left'>
	  <th>#</th>
	  <th>State</th>
	  <th>User</th>
	  <th>Title</th>
	</tr>
	{{range .Items}}
	<tr>
	  <td><a href='{{.HTMLURL}}'>{{.Number}}</td>
	  <td>{{.State}}</td>
	  <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
	  <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
	</tr>
	{{end}}
	</table>
	`

	var report = template.Must(template.New("issuelist").Parse(issueList))

	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	// run: go run issuehtml [search-term] > issues.html
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}
