package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type dollars float32

type database map[string]dollars

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

func (db database) list(w http.ResponseWriter, req *http.Request) {
	const skeleton = `
	<table>
	<tbody>
	<tr style='text-align: left'>
	<th>Item</th>
	<th>Price</th>
	</tr>
	{{range $item, $price := . }}
	<tr>
	<td>{{ $item }}</td>
	<td>{{ $price}}</td>
	</tr>
	</tbody>
	{{end}}
	</table>
	`
	var list = template.Must(template.New("store").Parse(skeleton))

	if err := list.Execute(w, db); err != nil {
		log.Fatal(err)
	}
}
