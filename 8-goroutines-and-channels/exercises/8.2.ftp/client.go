package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	handle(conn)
}

func handle(c net.Conn) {
	defer c.Close()

	in := bufio.NewScanner(os.Stdin)
	out := bufio.NewScanner(c)

	for in.Scan() {
		cmd := in.Text()

		if cmd == "" {
			continue
		}

		_, err := fmt.Fprintf(c, "%s\n", in.Text())
		if err != nil {
			return
		}

		out.Scan()
		rst := out.Text()

		if rst == "EOF" {
			return // close and disconnect
		}

		fmt.Printf("%s\n", rst)
	}
}
