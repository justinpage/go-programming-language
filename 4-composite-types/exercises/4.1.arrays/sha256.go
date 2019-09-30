package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Println(diff(&c1, &c2))
}

func diff(c1, c2 *[32]byte) int {
	var count int
	for i := range c1 {
		count += popCount(c1[i], c2[i])
	}
	return count
}

func popCount(b1, b2 *byte) int {
	var count int
	for i := uint(0); i < 8; i++ {
		// compare each bit against a truthy value
		bit1 := (b1 >> i) & 1
		bit2 := (b2 >> i) & 1

		if bit1 != bit2 {
			count++
		}
	}
	return count
}
