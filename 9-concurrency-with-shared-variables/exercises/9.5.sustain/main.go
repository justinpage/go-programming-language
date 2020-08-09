// Running this locally showed an average of 1,304,119 communications made over
// a period of ten seconds
package main

import (
	"fmt"
	"time"
)

func main() {
	ping := make(chan string)
	pong := make(chan string)

	go func() {
		for {
			ping <- "passed message"
		}
	}()

	var count int
	go func() {
		for {
			msg := <-ping
			count++ // count communication from ping
			pong <- fmt.Sprintf("%s: %d", msg, count)
		}
	}()

	tickerS := time.NewTicker(1 * time.Second)
	tickerM := time.NewTicker(60 * time.Second)
	var avg []int
	for {
		select {
		case <-tickerS.C:
			// first, capture total number of communications in one second
			avg = append(avg, count)
			// you could print the count here to show the result per second but
			// this might not show the average over longer periods of time.
			//
			// For example, what would the average be if count was captured each
			// second over a period of a minute? Would our average be the same
			// if more communications occurred in a minute, for a few seconds?
			//
			// The answer to this question is 1,313,265. Pretty close to what we
			// see per second but it is slightly higher. Still, having this as
			// proof for the average makes our solution more sound.
			count = 0
		case <-tickerM.C:
			// next, average the number of communications per second
			// over a period of ten seconds
			var sum int
			for _, v := range avg {
				sum += v
			}
			fmt.Printf("\raverage: %d", int(sum/len(avg)))
			avg = []int{}
		case <-pong:
			count++ // count communication from pong
		}
	}
}
