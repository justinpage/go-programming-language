// Provides a concurrency-safe bank with one account.
package main

import (
	"fmt"
)

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		}
	}
}

func init() {
	go teller()
}

func main() {
	Deposit(12)
	go Deposit(24)
	fmt.Println(Balance()) // 12
	fmt.Println(Balance()) // 36 --not guaranteed!
}
