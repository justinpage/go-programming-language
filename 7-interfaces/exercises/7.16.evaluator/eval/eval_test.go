package eval

import (
	"fmt"
	"testing"
)

func TestNewConcrete(t *testing.T) {
	tests := []struct {
		expr string
		env  Env
		want string
	}{
		{"x / y + [5, 2].length", Env{"x": 8, "y": 4}, "4"},
		{"pow(2, 3) + [5, 2, 1].min", Env{"F": 32}, "9"},
		{"(36 / 6) * [pow(x, y), sqrt(9), 5].max", Env{"x": 3, "y": 2}, "54"},
	}
	var prevExpr string
	for _, test := range tests {
		// Print expr only when it changes.
		if test.expr != prevExpr {
			fmt.Printf("\n%s\n", test.expr)
			prevExpr = test.expr
		}

		expr, err := Parse(test.expr)
		if err != nil {
			t.Error(err) // parse error
			continue
		}

		vars := make(map[Var]bool)
		if err := expr.Check(vars); err != nil {
			t.Error(err) // check error
			continue
		}

		expr, err = Parse(expr.String())
		if err != nil {
			t.Error(err) // parse error
			continue
		}

		got := fmt.Sprintf("%.6g", expr.Eval(test.env))
		fmt.Printf("\t%v => %s\n", test.env, got)
		if got != test.want {
			t.Errorf("%s.Eval() in %v = %q, want %q\n",
				test.expr, test.env, got, test.want)
		}
	}
}
