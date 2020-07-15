package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print("hello")
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

type client struct {
	address    string
	msg        chan string // an outgoing message channel
	idle       chan string // an idle message channel
	lastActive time.Time
}

var (
	entering = make(chan *client)
	leaving  = make(chan *client)
	messages = make(chan string) // all incoming client messages
)

func broadcaster() {
	clients := make(map[*client]bool) // all connected clients
	tick := time.Tick(10 * time.Second)
	for {
		select {
		case t := <-tick:
			for cli := range clients {
				if t.Sub(cli.lastActive).Seconds() > 10 {
					cli.idle <- "You are idle. Disconnecting."
				}
			}
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels
			for cli := range clients {
				cli.msg <- msg
			}
		case cli := <-entering:
			clients[cli] = true

			cli.msg <- "[Users]"
			for c := range clients {
				cli.msg <- fmt.Sprintf("[%s\t]", c.address)
			}
		case cli := <-leaving:
			delete(clients, cli)
			close(cli.msg)
			close(cli.idle)
		}
	}
}

func handleConn(conn net.Conn) {
	user := &client{
		address:    conn.RemoteAddr().String(),
		msg:        make(chan string), // outgoing client messages
		idle:       make(chan string),
		lastActive: time.Now(),
	}

	go clientWriter(conn, user)
	go clientCloser(conn, user)

	user.msg <- "You are " + user.address
	messages <- user.address + " has arrived"
	entering <- user

	input := bufio.NewScanner(conn)
	for input.Scan() {
		user.lastActive = time.Now()
		messages <- user.address + ": " + input.Text()
	}

	// NOTE: ignoring potential errors from input.Err()
	leaving <- user
	messages <- user.address + " has left"
}

func clientWriter(conn net.Conn, u *client) {
	for msg := range u.msg {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}

func clientCloser(conn net.Conn, u *client) {
	for msg := range u.idle {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
		conn.Close()
	}
}
