package main

import "fmt"

func join(sep string, vals ...string) string {
	if vals == nil {
		return ""
	}

	var response string
	for _, val := range vals {
		response += sep + val
	}
	return response[1:]
}

func main() {
	fmt.Println(join("-", "justin"))
	fmt.Println(join("-", "justin", "bob", "mike", "jeff"))

	values := []string{"justin", "bob", "mike", "jeff", " "}
	fmt.Println(join(" ", values...))
}
