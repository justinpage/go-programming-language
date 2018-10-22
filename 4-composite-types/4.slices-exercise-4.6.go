package main

import "fmt"

func main() {
	bytes := []byte("Hello, Justin")
	for _, b := range bytes {
		fmt.Printf("%x\n", b)
	}
}
