package main

import (
	"fmt"
)

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint64
}

func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// Return the number of elements
func (s *IntSet) Elems() (e []int) {
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				e = append(e, 64*i+j)
			}
		}
	}
	return
}

func main() {
	var x IntSet

	x.Add(1)
	x.Add(144)
	x.Add(9)

	for _, v := range x.Elems() {
		fmt.Println(v)
	}
}
