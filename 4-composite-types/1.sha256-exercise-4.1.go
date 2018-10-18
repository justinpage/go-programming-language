package main

import (
	"crypto/sha256"
	"fmt"
	"math/bits"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Println(Diff(&c1, &c2))
}

func Diff(c1, c2 *[32]uint8) int {
	result := (PopCount(c1) - PopCount(c2))
	if result < 0 {
		return -result
	}
	return result
}

func PopCount(c *[32]uint8) int {
	var count int
	for _, v := range c {
		count += bits.OnesCount8(v)
	}
	return count
}
