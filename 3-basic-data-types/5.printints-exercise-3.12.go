package main

import "fmt"

func main() {
	fmt.Println(isAnagram("abc", "cba"))
}

func isAnagram(s1, s2 string) bool {
	var counter1 rune
	var counter2 rune

	for _, v := range s1 {
		counter1 += v
	}

	for _, v := range s2 {
		counter2 += v
	}

	if counter1 == counter2 {
		return true
	}

	return false
}

func isAnagram2(s1, s2 string) bool {
	outline := make(map[string]bool)

	for _, v := range s1 {
		outline[string(v)] = true
	}

	for _, v := range s2 {
		if outline[string(v)] == false {
			return false
		}
	}

	return true
}
