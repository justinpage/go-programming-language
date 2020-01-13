// Clock1 is a TCP server that periodically writes the time.
package main

import (
	"flag"
	"io"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	tz := os.Getenv("TZ")
	if tz == "" {
		tz = "Local" // fallback to local time zone
	}

	port := flag.String("port", "8000", "port number")
	flag.Parse()

	listener, err := net.Listen("tcp", ":"+*port)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn, tz) // handle connections concurrently
	}
}

func handleConn(c net.Conn, tz string) {
	defer c.Close()

	location, err := time.LoadLocation(tz)
	if err != nil {
		return // e.g., invalid location
	}

	for {
		now := time.Now()
		_, err := io.WriteString(c, now.In(location).Format("15:04:05\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}
