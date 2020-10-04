// References:
// https://github.com/ray-g/gopl/blob/master/ch10/ex10.04/deps.go
//
// Command example: go run report.go strings strconv math math/rand
package main

import (
	"bufio"
	"bytes"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		os.Exit(0)
	}
	findTransitiveDependency(getImportPath(os.Args[1:]))
}

func getImportPath(p []string) []string {
	arg := []string{"list", "-f={{.ImportPath}}"}
	arg = append(arg, p...)
	out, err := exec.Command("go", arg...).Output()
	if err != nil {
		handleErr(err)
	}
	return strings.Fields(string(out))
}

func findTransitiveDependency(p []string) {
	target := make(map[string]bool)
	for _, v := range p {
		target[v] = true
	}
	arg := []string{"list",
		`-f={{.ImportPath}} {{join .Deps " "}}`,
		"github.com/justinpage/go-programming-language/10-packages-and-the-go-tool",
	}
	out, err := exec.Command("go", arg...).Output()
	if err != nil {
		handleErr(err)
	}
	input := bufio.NewScanner(bytes.NewReader(out))
	pkgs := make(map[string][]string)
	for input.Scan() {
		fields := strings.Fields(input.Text())
		pkg := fields[0]
		deps := fields[1:]
		for _, dep := range deps {
			if target[dep] {
				pkgs[pkg] = append(pkgs[pkg], dep)
			}
		}
	}
	for k, d := range pkgs {
		log.Printf("%s:\n", k)
		for _, v := range d {
			log.Printf("\t=> %s\n", v)
		}
	}
}

func handleErr(err error) {
	e, ok := err.(*exec.ExitError)
	if !ok {
		log.Fatalf("%s\n", err)
	}
	log.Fatalf("%s\n", e.Stderr)
}
