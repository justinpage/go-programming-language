package main

import (
	"html/template"
	"log"
	"os"
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

func printCities(cities []*City) {
	const cityList = `
	<table>
	<tr style='text-align: left'>
	  <th>Name</th>
	  <th>State</th>
	  <th>Order</th>
	</tr>
	{{range .}}
	<tr>
	  <td>{{.Name}}</td>
	  <td>{{.State}}</td>
	  <td>{{.Order}}</td>
	</tr>
	{{end}}
	</table>
	`

	var list = template.Must(template.New("citylist").Parse(cityList))

	if err := list.Execute(os.Stdout, cities); err != nil {
		log.Fatal(err)
	}
}

func (x stableSort) Len() int           { return len(x.t) }
func (x stableSort) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }
func (x stableSort) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }

func main() {
	listOfCities := cities()

	printCities(listOfCities)

	// byNameStable := func(x, y *CityOrder) bool {
	// 	if x.Name == y.Name {
	// 		return x.Order < y.Order
	// 	}
	// 	return x.Name < y.Name
	// }
	//
	// byStateStable := func(x, y *CityOrder) bool {
	// 	if x.State == y.State {
	// 		return x.Order < y.Order
	// 	}
	// 	return x.State < y.State
	// }
	//
	// var byActions []func(x, y *CityOrder) bool
	// byActions = append(byActions, byNameStable)
	// byActions = append(byActions, byStateStable)
	//
	// listOfCitiesWithOrder.Sort(byActions)
	//
	// printCitiesWithOrder(listOfCitiesWithOrder)
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
