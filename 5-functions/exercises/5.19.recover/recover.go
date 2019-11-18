package main

import "fmt"

func main() {
	fmt.Println(test())
}

func test() (val string) {
	type bailout struct{}

	defer func() {
		switch p := recover(); p {
		case bailout{}:
			val = "KaBlam!"
		default:
			panic(p)
		}
	}()

	panic(bailout{})
}
