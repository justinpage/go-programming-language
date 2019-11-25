package main

import (
	"bytes"
	"fmt"
)

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint64
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// IntersectWith sets s to the insersect values of t.
func (s *IntSet) IntersectWith(t *IntSet) {
	for i := range t.words {
		for j := range s.words {
			if i == j {
				s.words[j] &= t.words[i]
			}
		}
	}
}

// DifferenceWith sets s to the difference of s and t.
func (s *IntSet) DifferenceWith(t *IntSet) {
	for i := range t.words {
		for j := range s.words {
			if i == j {
				s.words[j] &^= t.words[i]
			}
		}
	}
}

// SymetricDifference sets s to the symetric difference of s and t.
func (s *IntSet) SymetricDifference(t *IntSet) {
	for i := range t.words {
		for j := range s.words {
			if i == j {
				s.words[j] ^= t.words[i]
			}
		}
	}
}

func (s *IntSet) Clear() {
	s.words = nil
}

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func main() {
	var x, y IntSet

	// IntersectWith
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(3)

	y.Add(2)
	y.Add(8)
	y.Add(144)
	y.Add(9)
	y.Add(4)

	x.IntersectWith(&y)
	fmt.Println(x.String())

	x.Clear()
	y.Clear()

	// DifferenceWith
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(3)

	y.Add(2)
	y.Add(8)
	y.Add(144)
	y.Add(9)
	y.Add(4)

	x.DifferenceWith(&y)
	fmt.Println(x.String())

	x.Clear()
	y.Clear()

	// SymetricDifference
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(3)

	y.Add(2)
	y.Add(8)
	y.Add(144)
	y.Add(9)
	y.Add(4)

	x.SymetricDifference(&y)
	fmt.Println(x.String())
}
