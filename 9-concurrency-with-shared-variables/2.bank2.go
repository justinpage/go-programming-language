// Provides a concurrency-safe bank with one account.
package main

import (
	"fmt"
)

var (
	sema    = make(chan struct{}, 1)
	balance int
)

func Deposit(amount int) {
	sema <- struct{}{} // acquire token
	balance = balance + amount
	<-sema // release token
}

func Balance() int {
	sema <- struct{}{} // acquire token
	b := balance
	<-sema // release token
	return b
}

func main() {
	Deposit(12)
	go Deposit(24)
	fmt.Println(Balance()) // 12
	fmt.Println(Balance()) // 36
}
