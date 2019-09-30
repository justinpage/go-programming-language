package main

import "fmt"

func main() {
	data := []string{"one", "three", "three"}
	fmt.Println(unique(data))
}

func unique(data []string) []string {
	var output []string

	for _, v := range data {
		if !has(output, v) {
			output = append(output, v)
		}
	}
	return output
}

func has(data []string, t string) bool {
	for _, v := range data {
		if v == t {
			return true
		}
	}
	return false
}
