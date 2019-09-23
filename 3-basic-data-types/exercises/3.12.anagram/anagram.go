package main

import (
	"fmt"
)

func main() {
	fmt.Println(isAnagram("binary", "brainy"))
	fmt.Println(isAnagram("listen", "silent"))
	fmt.Println(isAnagram("brainy", "listen"))
}

func isAnagram(x, y string) bool {
	m := make(map[rune]bool)

	for _, v := range x {
		m[v] = true
	}

	for _, v := range y {
		if m[v] == false {
			return false
		}
	}

	return true
}
