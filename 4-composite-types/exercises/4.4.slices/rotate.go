package main

import "fmt"

func rotateLeft(s []int, n int) {
	var stack []int

	stack = append(stack, s[n:]...)
	stack = append(stack, s[:n]...)

	copy(s, stack)
}

func main() {
	a := []int{0, 1, 2, 3, 4, 5}
	rotateLeft(a, 2)
	fmt.Println(a)
}
