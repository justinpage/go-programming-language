package main

import "fmt"

func main() {
	var graph = make(map[string]map[string]bool)

	addEdge := func(from, to string) {
		edges := graph[from]
		if edges == nil {
			edges = make(map[string]bool)
			graph[from] = edges
		}
		edges[to] = true
	}

	hasEdge := func(from, to string) bool {
		return graph[from][to]
	}

	addEdge("hello", "i am here")

	fmt.Println(hasEdge("hello", "i am here"))
	fmt.Println(hasEdge("hello", "i am not here"))

	fmt.Println(graph)
}
