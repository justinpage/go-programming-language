package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	for _, arg := range os.Args[1:] {
		z := strings.Split(arg, "=")
		go dial(z[0], z[1])
	}
	select {} // block without eating CPU
}

func dial(name, port string) {
	conn, err := net.Dial("tcp", ":"+port)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(name, conn)
}

func mustCopy(n string, src io.Reader) {
	input := bufio.NewScanner(src)
	for input.Scan() {
		fmt.Printf("%-10s%8s\n", n, input.Text())
	}
}
