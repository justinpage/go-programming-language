package echo

import "testing"

func BenchmarkRange(b *testing.B) {
	args := []string{"hello", "range"}
	for i := 0; i < b.N; i++ {
		Range(args)
	}
}

func BenchmarkLoop(b *testing.B) {
	args := []string{"hello", "loop"}
	for i := 0; i < b.N; i++ {
		Loop(args)
	}
}

func BenchmarkJoin(b *testing.B) {
	args := []string{"hello", "join"}
	for i := 0; i < b.N; i++ {
		Join(args)
	}
}
