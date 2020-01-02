package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/justinpage/go-programming-language/7-interfaces/exercises/7.16.evaluator/eval"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", http.StripPrefix("/", fs))

	http.HandleFunc("/calculate", calculate)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func calculate(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	expr, err := parseAndCheck(r.Form.Get("expr"))
	if err != nil {
		http.Error(w, "bad expr: "+err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "%.6g", expr.Eval(eval.Env{}))
}

func parseAndCheck(s string) (eval.Expr, error) {
	if s == "" {
		return nil, fmt.Errorf("empty expression")
	}
	expr, err := eval.Parse(s)
	if err != nil {
		return nil, err
	}
	vars := make(map[eval.Var]bool)
	if err := expr.Check(vars); err != nil {
		return nil, err
	}
	return expr, nil
}
