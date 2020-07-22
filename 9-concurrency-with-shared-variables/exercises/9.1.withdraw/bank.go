// Provides a concurrency-safe bank with one account.
package main

import (
	"fmt"
)

type withdrawal struct {
	amount int       // amount to withdraw
	ch     chan bool // result of withdrawal
}

var deposits = make(chan int)           // send amount to deposit
var balances = make(chan int)           // receive balance
var withdrawals = make(chan withdrawal) // withdraw amount from balance

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }

func Withdraw(amount int) bool {
	var w = withdrawal{amount, make(chan bool)}
	withdrawals <- w
	return <-w.ch
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case withdraw := <-withdrawals:
			if withdraw.amount > balance {
				withdraw.ch <- false
			}
			balance -= withdraw.amount
			withdraw.ch <- true
		case balances <- balance:
		}
	}
}

func init() {
	go teller()
}

func main() {
	Deposit(12)
	fmt.Println(Balance())    // 12
	fmt.Println(Withdraw(6))  // true
	fmt.Println(Balance())    // 6
	fmt.Println(Withdraw(12)) // false
}
