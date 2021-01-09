package intset

import (
	"bytes"
	"fmt"
	"sort"
	"testing"
)

type hash map[int]bool

func TestIntSet(t *testing.T) {
	var x, y IntSet
	var want hash = map[int]bool{1: true, 9: true}
	// Add
	x.Add(1)
	x.Add(9)
	if got := x.String(); got != want.String() {
		t.Errorf(`IntSet.Add(%v) = %v`, want, got)
	}
	// Has
	if got := x.Has(1); got != true {
		t.Errorf(`IntSet.Has(%v) = %v`, 1, got)
	}
	if got := x.Has(144); got != false {
		t.Errorf(`IntSet.Has(%v) = %v`, 144, got)
	}
	// UnionWith
	y.Add(9)
	y.Add(42)
	x.UnionWith(&y)
	want[42] = true
	if got := x.String(); got != want.String() {
		t.Errorf(`IntSet.UnionWith(%v) = %v`, want, got)
	}
	// String
	if got := x.String(); got != want.String() {
		t.Errorf(`IntSet.String(%v) = %v`, want, got)
	}
}

func (h hash) String() string {
	keys := make([]int, 0)
	for k := range h {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	var buf bytes.Buffer
	buf.WriteByte('{')
	for _, k := range keys {
		if buf.Len() > len("{") {
			buf.WriteByte(' ')
		}
		fmt.Fprintf(&buf, "%d", k)
	}
	buf.WriteByte('}')
	return buf.String()
}
