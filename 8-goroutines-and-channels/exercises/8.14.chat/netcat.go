package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{}) // NOTE: ignoring errors
	go func() {
		io.Copy(os.Stdout, conn)
		conn.(*net.TCPConn).CloseRead()
		done <- struct{}{} // signal the main goroutine
	}()
	go func() {
		mustCopy(conn, os.Stdin)
		conn.(*net.TCPConn).CloseWrite()
	}()
	<-done // wait for background goroutine to finish
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
