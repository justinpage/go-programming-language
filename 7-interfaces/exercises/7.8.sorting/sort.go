package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
)

type City struct {
	Name  string
	State string
}

type CityOrder struct {
	City
	Order int
}

type byName []*City
type byState []*City

type byCityOrder []*CityOrder

type stableSort struct {
	t    []*CityOrder
	less func(x, y *CityOrder) bool
}

func cities() []*City {
	return []*City{
		{"Chicago", "IL"}, {"Champaign", "IL"}, {"Detroit", "MI"},
		{"New York", "NY"}, {"Buffalo", "NY"}, {"Milwaukee", "WI"},
		{"Albany", "NY"}, {"Green Bay", "WI"}, {"Syracuse", "NY"},
		{"Rockford", "IL"}, {"Evanston", "IL"},
	}
}

func citiesWithOrder() byCityOrder {
	return byCityOrder{
		{City{"Chicago", "IL"}, 0}, {City{"Champaign", "IL"}, 0},
		{City{"Detroit", "MI"}, 0}, {City{"New York", "NY"}, 0},
		{City{"Buffalo", "NY"}, 0}, {City{"Milwaukee", "WI"}, 0},
		{City{"Albany", "NY"}, 0}, {City{"Green Bay", "WI"}, 0},
		{City{"Syracuse", "NY"}, 0}, {City{"Rockford", "IL"}, 0},
		{City{"Evanston", "IL"}, 0},
	}
}

func printCities(cities []*City) {
	const format = "%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Name", "State")
	fmt.Fprintf(tw, format, "----", "-----")
	for _, c := range cities {
		fmt.Fprintf(tw, format, c.Name, c.State)
	}
	tw.Flush() // calculate column widths and print table
}

func printCitiesWithOrder(cities []*CityOrder) {
	const format = "%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Name", "State", "Order")
	fmt.Fprintf(tw, format, "----", "-----", "-----")
	for _, c := range cities {
		fmt.Fprintf(tw, format, c.Name, c.State, c.Order)
	}
	tw.Flush() // calculate column widths and print table
}

func (x byName) Len() int           { return len(x) }
func (x byName) Less(i, j int) bool { return x[i].Name < x[j].Name }
func (x byName) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func (x byState) Len() int           { return len(x) }
func (x byState) Less(i, j int) bool { return x[i].State < x[j].State }
func (x byState) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func (x stableSort) Len() int           { return len(x.t) }
func (x stableSort) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }
func (x stableSort) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }

func main() {
	// Stability not guaranteed
	listOfCities := cities()

	printCities(listOfCities)

	sort.Stable(byName(listOfCities))
	sort.Stable(byState(listOfCities))

	printCities(listOfCities)

	// Explicit stability
	listOfCitiesWithOrder := citiesWithOrder()

	byNameStable := func(x, y *CityOrder) bool {
		if x.Name == y.Name {
			return x.Order < y.Order
		}
		return x.Name < y.Name
	}

	byStateStable := func(x, y *CityOrder) bool {
		if x.State == y.State {
			return x.Order < y.Order
		}
		return x.State < y.State
	}

	var byActions []func(x, y *CityOrder) bool
	byActions = append(byActions, byNameStable)
	byActions = append(byActions, byStateStable)

	listOfCitiesWithOrder.Sort(byActions)

	printCitiesWithOrder(listOfCitiesWithOrder)
}

func (c byCityOrder) Sort(byActions []func(x, y *CityOrder) bool) {
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
