// Provides a concurrency-safe bank with one account.
package main

import (
	"fmt"
	"sync"
)

var (
	mu      sync.Mutex // guards balance
	balance int
)

func Deposit(amount int) {
	mu.Lock()
	balance = balance + amount
	defer mu.Unlock()
}

func Balance() int {
	mu.Lock()
	defer mu.Unlock()
	return balance
}

func main() {
	Deposit(12)
	go Deposit(24)
	fmt.Println(Balance()) // 12
	fmt.Println(Balance()) // 36
}
