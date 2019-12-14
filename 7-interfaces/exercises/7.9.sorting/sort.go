package main

import (
	"html/template"
	"log"
	"net/http"
	"sort"
)

type City struct {
	Name  string
	State string
	Order int
}

type byCityOrder []*City

type stableSort struct {
	t    []*City
	less func(x, y *City) bool
}

func cities() byCityOrder {
	return byCityOrder{
		{"Chicago", "IL", 0}, {"Champaign", "IL", 0},
		{"Detroit", "MI", 0}, {"New York", "NY", 0},
		{"Buffalo", "NY", 0}, {"Milwaukee", "WI", 0},
		{"Albany", "NY", 0}, {"Green Bay", "WI", 0},
		{"Syracuse", "NY", 0}, {"Rockford", "IL", 0},
		{"Evanston", "IL", 0},
	}
}

func (x stableSort) Len() int           { return len(x.t) }
func (x stableSort) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }
func (x stableSort) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }

func byNameStable(x, y *City) bool {
	if x.Name == y.Name {
		return x.Order < y.Order
	}
	return x.Name < y.Name
}

func byStateStable(x, y *City) bool {
	if x.State == y.State {
		return x.Order < y.Order
	}
	return x.State < y.State
}

func main() {
	listOfCities := cities()

	var byActions []func(x, y *City) bool

	printCities := func(w http.ResponseWriter, r *http.Request) {
		const cityList = `
		<style> th { cursor: pointer } </style>
		<table>
		<tbody>
		<tr style='text-align: left'>
		  <th onclick="sortBy('name');">Name</th>
		  <th onclick="sortBy('state');">State</th>
		  <th>Order</th>
		</tr>
		{{range .}}
		<tr>
		  <td>{{.Name}}</td>
		  <td>{{.State}}</td>
		  <td>{{.Order}}</td>
		</tr>
		</tbody>
		{{end}}
		</table>
		<script>
			async function sortBy(column) {
				let table = document.getElementsByTagName("table")[0]
				const actions = {
					"name": "order-by-name",
					"state": "order-by-state"
				}

				response = await fetch(actions[column])
				table.innerHTML = await response.text()
			}
		</script>
		`

		var list = template.Must(template.New("citylist").Parse(cityList))

		if err := list.Execute(w, listOfCities); err != nil {
			log.Fatal(err)
		}
	}

	orderByName := func(w http.ResponseWriter, r *http.Request) {
		byActions = append(byActions, byNameStable)
		listOfCities.Sort(byActions)
		printCities(w, r)
	}

	orderByState := func(w http.ResponseWriter, r *http.Request) {
		byActions = append(byActions, byStateStable)
		listOfCities.Sort(byActions)
		printCities(w, r)
	}

	http.HandleFunc("/", printCities)
	http.HandleFunc("/order-by-name", orderByName)
	http.HandleFunc("/order-by-state", orderByState)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func (c byCityOrder) Sort(byActions []func(x, y *City) bool) {
	sort.Sort(stableSort{c, byActions[0]})

	// Keep order from first sort
	for i, v := range c {
		v.Order = i
	}

	if len(byActions) > 1 {
		for i := 1; i < len(byActions); i++ {
			sort.Sort(stableSort{c, byActions[i]})
		}
	}
}
