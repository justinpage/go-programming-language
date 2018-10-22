package main

import "fmt"

func main() {
	data := []string{"one", "three", "three"}
	uniq := Unique(data)

	fmt.Printf("%q\n", uniq)
}

func Unique(collection []string) []string {
	set := make([]string, 0)

	for _, v := range collection {
		if !Has(set, v) {
			set = append(set, v)
		}
	}

	return set
}

func Location(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}

	return -1
}

func Has(vs []string, t string) bool {
	return Location(vs, t) >= 0
}
