package main

import "fmt"

func max(vals ...int) int {
	total := 0
	for _, val := range vals {
		if val > total {
			total = val
		}
	}
	return total
}

func min(vals ...int) int {
	if vals == nil {
		return 0
	}

	total := vals[0]
	for _, val := range vals {
		if val < total {
			total = val
		}
	}
	return total
}

func main() {
	fmt.Println(max())
	fmt.Println(max(3))
	fmt.Println(max(1, 2, 3, 4))

	values := []int{1, 2, 3, 4}
	fmt.Println(max(values...))

	fmt.Println(min())
	fmt.Println(min(3))
	fmt.Println(min(1, 2, 3, 4))

	fmt.Println(min(values...))
}
