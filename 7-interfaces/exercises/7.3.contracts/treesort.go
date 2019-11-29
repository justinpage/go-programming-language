package main

import (
	"bytes"
	"fmt"
)

type tree struct {
	value       int
	left, right *tree
}

func main() {
	list := []int{2, 1, 4, 3, 5}
	Sort(list)
	fmt.Println(list)
}

// Sort sorts values in place.
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	fmt.Println(root)
	appendValues(values[:0], root)
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func (t *tree) String() string {
	var buf bytes.Buffer
	var traverse func(t *tree)

	traverse = func(t *tree) {
		if t != nil {
			if buf.Len() > 0 {
				buf.WriteByte(' ')
			}
			fmt.Fprintf(&buf, "%d", t.value)
			traverse(t.left)
			traverse(t.right)
		}
	}

	traverse(t)

	return buf.String()
}
