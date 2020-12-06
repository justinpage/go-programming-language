package charcount

import "testing"

func TestCharcount(t *testing.T) {
	var tests = []struct {
		input string
		want  error
	}{
		{"a", nil},
		{"aa", nil},
		{"kayak", nil},
		{"detartrated", nil},
		{"A man, a plan, a canal: Panama", nil},
		{"Evil I did dwell; lewd did I live.", nil},
		{"Able was I ere I saw Elba", nil},
		{"été", nil},
		{"Et se resservir, ivresse reste.", nil},
	}
	for _, test := range tests {
		if got := CharCount(test.input); got != test.want {
			t.Errorf(`CharCount(%q) = %s`, test.input, got)
		}
	}
}
