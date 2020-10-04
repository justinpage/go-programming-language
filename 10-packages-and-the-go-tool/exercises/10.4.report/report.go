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
	for _, p := range os.Args[1:] {
		findTransitiveDependency(getImportPath(p))
	}
}

func getImportPath(p string) (a string) {
	arg := []string{"list", "-f={{.ImportPath}}", p}
	out, err := exec.Command("go", arg...).Output()
	if err != nil {
		handleErr(err)
		return
	}
	return strings.TrimSuffix(string(out), "\n")

}

func findTransitiveDependency(p string) {
	arg := []string{"list",
		`-f={{.ImportPath}} {{join .Deps " "}}`,
		"github.com/justinpage/go-programming-language/10-packages-and-the-go-tool",
	}
	out, err := exec.Command("go", arg...).Output()
	if err != nil {
		handleErr(err)
	}
	input := bufio.NewScanner(bytes.NewReader(out))
	pkgs := make(map[string]string)
	for input.Scan() {
		fields := strings.Fields(input.Text())
		pkg := fields[0]
		deps := fields[1:]
		for _, dep := range deps {
			if p == dep {
				pkgs[pkg] = p
			}
		}
	}
	for k, v := range pkgs {
		log.Printf("%s -> %v", k, v)
	}
}

func handleErr(err error) {
	e, ok := err.(*exec.ExitError)
	if !ok {
		log.Printf("%s\n", err)
	}
	log.Printf("%s\n", e.Stderr)
}
