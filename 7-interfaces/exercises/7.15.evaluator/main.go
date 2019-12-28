package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/justinpage/go-programming-language/7-interfaces/exercises/7.15.evaluator/eval"
)

func main() {
	for {
		fmt.Print("¬ : ")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()

		input := scanner.Text()

		if input == "" {
			fmt.Fprintln(os.Stderr, "--¬ Expression required (e.g. 2 + 2)")
			continue
		}

		expr, err := eval.Parse(input)
		if err != nil {
			fmt.Fprintln(os.Stderr, "--¬ Invalid expression. Please try again.")
			continue
		}

		vars := make(map[eval.Var]bool)

		if err := expr.Check(vars); err != nil {
			fmt.Fprintln(os.Stderr, "--¬ Invalid expression. Please try again.")
			continue
		}

		env := environment(vars)

		fmt.Fprintf(os.Stdout, "¬ = %.6g\n", expr.Eval(env))
	}
}

func environment(vars map[eval.Var]bool) eval.Env {
	env := make(eval.Env)
	scanner := bufio.NewScanner(os.Stdin)

	for k := range vars {
		for {
			fmt.Printf("¬ : %s = ", k)
			scanner.Scan()
			v, err := strconv.ParseFloat(scanner.Text(), 64)
			if err != nil {
				fmt.Fprintln(os.Stderr, "--¬ Invalid value. Try again.")
				continue
			}
			env[k] = v
			break
		}
	}

	return env
}
