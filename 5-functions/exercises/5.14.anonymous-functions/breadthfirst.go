package main

import (
	"fmt"
)

// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

// breadthFirst calls f for each item in the worklist.
// Any items returned by f are added to the worklist.
// f is called at most once for each item.
func breadthFirst(m map[string][]string, worklist []string) []string {
	seen := make(map[string]bool)
	var order []string

	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, m[item]...)
				order = append(order, item)
			}
		}
	}

	return order
}

func main() {
	var keys []string
	for key := range prereqs {
		keys = append(keys, key)
	}

	for i, course := range breadthFirst(prereqs, keys) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}
