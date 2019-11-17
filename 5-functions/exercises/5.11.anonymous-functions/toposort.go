package main

import (
	"fmt"
	"sort"
)

// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string][]string{
	"algorithms":     {"data structures"},
	"calculus":       {"linear algebra"},
	"linear algebra": {"calculus"}, // cycle

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"intro to programming":  {"data structures"}, // cycle
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) []string {
	var order []string

	seen := make(map[string]bool)
	cycle := make(map[string][]string)

	var visitAll func(items []string)
	var findCircle func(item string)

	endlessCycle := func(vs map[string]int) bool {
		for _, v := range vs {
			if v > 1 {
				return true
			}
		}
		return false
	}

	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			} else if cycle[item] == nil {
				findCircle(item)
			}
		}
	}

	findCircle = func(item string) {
		var find func(items string)
		var courses []string

		find = func(node string) {
			itemCount := make(map[string]int)

			for _, course := range courses {
				itemCount[course]++
			}

			// Found cycle through dependencies
			if endlessCycle(itemCount) {
				cycle[item] = courses
				return
			}

			courses = append(courses, node)

			for _, dep := range m[node] {
				// Found cycle directly
				if dep == item {
					// Append item again to support console formatting
					courses = append(courses, item)
					cycle[item] = courses
					return
				}
				// Dig deeper
				find(dep)
			}
		}

		// Begin cycle search
		find(item)
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	visitAll(keys)

	for course, _ := range cycle {
		fmt.Printf("Cycle found for %s:\n", course)

		i := 0
		for _, preq := range cycle[course] {
			fmt.Printf("%*s%s->\n", i*2, "", preq)
			i++
		}
		fmt.Println()
	}

	return order
}
