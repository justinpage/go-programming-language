package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
)

const (
	ServiceReadyForNewUser = "220 Service ready for new user\n"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn) // handle connections concurrently
	}
}

func handleConn(c net.Conn) {
	defer c.Close()

	_, err := io.WriteString(c, ServiceReadyForNewUser)
	if err != nil {
		return // e.g., client disconnected
	}

	cmd := bufio.NewScanner(c)
	for cmd.Scan() {
		switch cmd := cmd.Text(); cmd {
		case "ls":
			fmt.Fprint(c, "list of items\n")
		case "get":
			fmt.Fprint(c, "here is an item\n")
		case "close":
			fmt.Fprint(c, errors.New("EOF").Error())
			return
		default:
			fmt.Fprintf(c, "command not found: %s\n", cmd)
		}
	}
}
