package split

import (
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	var tests = []struct {
		input string
		sep   string
		want  int
	}{
		{"a:b:c", ":", 3},
		{"a b c", " ", 3},
		{"abc", "", 3},
	}
	for _, test := range tests {
		words := strings.Split(test.input, test.sep)
		if got, want := len(words), test.want; got != want {
			t.Errorf("Split(%q, %q) returned %d words, want %d",
				test.input, test.sep, got, want)
		}
	}
}
