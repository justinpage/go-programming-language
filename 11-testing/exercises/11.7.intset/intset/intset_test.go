package intset

import (
	"bytes"
	"fmt"
	"math/rand"
	"sort"
	"testing"
)

const SEED = 1611961022311201000

type hash map[int]bool

func (x hash) UnionWith(y hash) {
	for k, v := range y {
		x[k] = v
	}
}

func (h hash) String() string {
	keys := make([]int, len(h))
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

func BenchmarkAddIntSet10(b *testing.B) {
	benchmarkAddIntSet(b, 10)
}

func BenchmarkAddIntSet100(b *testing.B) {
	benchmarkAddIntSet(b, 100)
}

func BenchmarkAddIntSet1000(b *testing.B) {
	benchmarkAddIntSet(b, 1000)
}

func benchmarkAddIntSet(b *testing.B, size int) {
	rng := rand.New(rand.NewSource(SEED))
	for i := 0; i < size; i++ {
		var x IntSet
		x.Add(randomInt(rng, size))
		x.Add(randomInt(rng, size))
	}
}

func BenchmarkUnionWithIntSet10(b *testing.B) {
	benchmarkUnionWithIntSet(b, 10)
}

func BenchmarkUnionWithIntSet100(b *testing.B) {
	benchmarkUnionWithIntSet(b, 100)
}

func BenchmarkUnionWithIntSet1000(b *testing.B) {
	benchmarkUnionWithIntSet(b, 1000)
}

func benchmarkUnionWithIntSet(b *testing.B, size int) {
	rng := rand.New(rand.NewSource(SEED))
	for i := 0; i < size; i++ {
		var x, y IntSet
		x.Add(randomInt(rng, size))
		x.Add(randomInt(rng, size))
		y.Add(randomInt(rng, size))
		y.Add(randomInt(rng, size))
		x.UnionWith(&y)
	}
}

func BenchmarkStringIntSet10(b *testing.B) {
	benchmarkStringIntSet(b, 10)
}

func BenchmarkStringIntSet100(b *testing.B) {
	benchmarkStringIntSet(b, 100)
}

func BenchmarkStringIntSet1000(b *testing.B) {
	benchmarkStringIntSet(b, 1000)
}

func benchmarkStringIntSet(b *testing.B, size int) {
	rng := rand.New(rand.NewSource(SEED))
	for i := 0; i < size; i++ {
		var x IntSet
		x.Add(randomInt(rng, size))
		x.Add(randomInt(rng, size))
		x.String()
	}
}

func BenchmarkAddMap10(b *testing.B) {
	benchmarkAddMap(b, 10)
}

func BenchmarkAddMap100(b *testing.B) {
	benchmarkAddMap(b, 100)
}

func BenchmarkAddMap1000(b *testing.B) {
	benchmarkAddMap(b, 1000)
}

func benchmarkAddMap(b *testing.B, size int) {
	rng := rand.New(rand.NewSource(SEED))
	for i := 0; i < size; i++ {
		x := make(hash)
		x[randomInt(rng, size)] = true
		x[randomInt(rng, size)] = true
	}
}

func BenchmarkUnionWithMap10(b *testing.B) {
	benchmarkUnionWithMap(b, 10)
}

func BenchmarkUnionWithMap100(b *testing.B) {
	benchmarkUnionWithMap(b, 100)
}

func BenchmarkUnionWithMap1000(b *testing.B) {
	benchmarkUnionWithMap(b, 1000)
}

func benchmarkUnionWithMap(b *testing.B, size int) {
	rng := rand.New(rand.NewSource(SEED))
	for i := 0; i < size; i++ {
		x := hash{randomInt(rng, size): true, randomInt(rng, size): true}
		y := hash{randomInt(rng, size): true, randomInt(rng, size): true}
		x.UnionWith(y)
	}
}

func BenchmarkStringMap10(b *testing.B) {
	benchmarkStringMap(b, 10)
}

func BenchmarkStringMap100(b *testing.B) {
	benchmarkStringMap(b, 100)
}

func BenchmarkStringMap1000(b *testing.B) {
	benchmarkStringMap(b, 1000)
}

func benchmarkStringMap(b *testing.B, size int) {
	rng := rand.New(rand.NewSource(SEED))
	for i := 0; i < size; i++ {
		x := make(hash)
		x[randomInt(rng, size)] = true
		x[randomInt(rng, size)] = true
		x.String()
	}
}

func randomInt(rng *rand.Rand, length int) int {
	return rng.Intn(length)
}
