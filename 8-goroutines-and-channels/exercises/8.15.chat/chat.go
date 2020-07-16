package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
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
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

type client struct {
	name       string
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
	tick := time.Tick(60 * time.Second)
	for {
		select {
		case t := <-tick:
			for cli := range clients {
				if t.Sub(cli.lastActive).Seconds() > 60 {
					cli.idle <- "You are idle. Disconnecting."
				}
			}
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels
			for cli := range clients {
				// example of non-blocking send
				select {
				case cli.msg <- msg:
					log.Printf("Sent message to %s\n", cli.name)
				default:
					log.Printf("No message to %s\n", cli.name)
				}
			}
		case cli := <-entering:
			clients[cli] = true

			cli.msg <- "[Users]"
			for c := range clients {
				cli.msg <- fmt.Sprintf("[%s\t]", c.name)
			}
		case cli := <-leaving:
			delete(clients, cli)
			close(cli.msg)
			close(cli.idle)
		}
	}
}

func handleConn(conn net.Conn) {
	rd := bufio.NewReader(conn)
	fmt.Fprint(conn, "Enter name: ")
	name, err := rd.ReadString('\n')
	if err != nil {
		log.Println(err)
		log.Println("Using remote network address as default name")
	}
	name = strings.TrimSpace(name)

	user := &client{
		name:       name,
		msg:        make(chan string), // outgoing client messages
		idle:       make(chan string),
		lastActive: time.Now(),
	}

	go clientWriter(conn, user)
	go clientCloser(conn, user)

	user.msg <- "You are " + user.name
	messages <- user.name + " has arrived"
	entering <- user

	input := bufio.NewScanner(conn)
	for input.Scan() {
		user.lastActive = time.Now()
		messages <- user.name + ": " + input.Text()
	}

	// NOTE: ignoring potential errors from input.Err()
	leaving <- user
	messages <- user.name + " has left"
}

func clientWriter(conn net.Conn, u *client) {
	delay := time.NewTicker(2 * time.Second)
	for msg := range u.msg {
		<-delay.C               // adding delay to simulate blocking receive
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
	delay.Stop()
}

func clientCloser(conn net.Conn, u *client) {
	for msg := range u.idle {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
		conn.Close()
	}
}
