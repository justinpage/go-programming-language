package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	go clear()
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

func clear() {
	for {
		time.Sleep(1 * time.Second)
		print("\033[H\033[2J") // https://stackoverflow.com/a/22892171/2395590
	}
}
