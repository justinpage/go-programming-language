package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	// log.SetPrefix("wait: ") // attractive output
	// log.SetFlags(0) // suppress the display of date and time

	for _, url := range os.Args[1:] {
		if err := WaitForServer(url); err != nil {
			log.Printf("Site is down: %v\n", err)
		}
	}
}

// WaitForServer attempts to contact the server of a URL.
// It tried for one minute using exponential back-off.
// It reports an error if all attempts fail.
func WaitForServer(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil // success
		}
		log.Printf("server not responding (%s); retrying...", err)
		time.Sleep(time.Second << uint(tries)) // exponential back-off
	}
	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}
