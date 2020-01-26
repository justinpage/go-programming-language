// References:
// https://cr.yp.to/ftp.html
// https://github.com/torbiak/gopl/tree/master/ex8.2
package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

// List of FTP server return codes
const (
	AcceptedDataConnection      = "150 Accepted data connection\n"
	SystemStatus                = "211 no-features\n"
	NameSystemType              = "215 UNIX Type: L8\n"
	ServiceReadyForNewUser      = "220 Service ready for new user\n"
	ServiceClosingConnection    = "221 Service closing control connection\n"
	ClosingDataConnection       = "226 Closing data connection\n"
	EnteringPassiveMode         = "227 Entering Passive Mode (%s)\n"
	UserLoggedInProceed         = "230 Logged in %s, proceed\n"
	PathNameCreated             = "257 Created \"%s\"\n"
	CurrentWorkingDirectory     = "257 \"%s\"\n"
	RequestedFileActionNotTaken = "450 Requested file action not taken\n"
	RequestedActionHasFailed    = "500 Requested action has failed \"%s\"\n"
	CommandNotImplemented       = "502 Command not implemented \"%s\"\n"
)

type server struct {
	conn net.Conn
	pasv net.Listener
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}

		s := &server{conn: conn}
		s.handleResponse(ServiceReadyForNewUser) // automatically accept

		go handleConn(s) // handle connections concurrently
	}
}

func handleConn(s *server) {
	defer s.conn.Close()

	err := os.Chdir("/tmp/") // start each connection inside /tmp/ dir
	if err != nil {
		return // unable to use /tmp/ directory
	}

	temp, err := ioutil.TempDir("/tmp/", "ftp-")
	if err != nil {
		return // unable to create temporary directory
	}

	defer os.RemoveAll(temp)

	cmd := bufio.NewScanner(s.conn)
	for cmd.Scan() {
		cmd := cmd.Text()
		arg := strings.Split(cmd, " ")

		if len(arg) > 1 {
			cmd, arg = arg[0], arg[1:]
		}

		switch cmd {
		case "USER":
			s.handleResponse(fmt.Sprintf(UserLoggedInProceed, arg[0]))
		case "SYST":
			s.handleResponse(fmt.Sprintf(NameSystemType))
		case "FEAT":
			s.handleResponse(SystemStatus)
		case "QUIT":
			s.handleResponse(ServiceClosingConnection)
			return
		case "EPSV":
			s.handleResponse(fmt.Sprintf(CommandNotImplemented, cmd))
		case "PASV":
			s.handlePASV()
		case "LIST":
			s.handleList(arg)
		case "PWD":
			cwd, _ := os.Getwd()
			s.handleResponse(fmt.Sprintf(CurrentWorkingDirectory, cwd))
		default:
			fmt.Println("cmd", cmd)
			s.handleResponse(fmt.Sprintf(CommandNotImplemented, cmd))
		}
	}
}

func (s *server) handleResponse(msg string) {
	_, err := io.WriteString(s.conn, msg)
	if err != nil {
		return // e.g., client disconnected
	}
}

func (s *server) handlePASV() {
	var err error
	s.pasv, err = net.Listen("tcp", "") // port automatically chosen

	_, p, err := net.SplitHostPort(s.pasv.Addr().String())
	h, _, err := net.SplitHostPort(s.conn.LocalAddr().String())

	addr, err := net.ResolveIPAddr("", h)
	port, err := strconv.ParseInt(p, 10, 64)

	ip := addr.IP.To4()

	location := fmt.Sprintf(
		"%d,%d,%d,%d,%d,%d", ip[0], ip[1], ip[2], ip[3], port/256, port%256,
	)

	if err != nil {
		log.Print(err)
		s.handleResponse(fmt.Sprintf(RequestedActionHasFailed, "PASV"))
	}

	s.handleResponse(fmt.Sprintf(EnteringPassiveMode, location))
}

func (s *server) handleList(arg []string) {
	conn, err := s.pasv.Accept()
	if err != nil {
		log.Print(err) // e.g., connection aborted
		s.handleResponse(fmt.Sprintf(RequestedActionHasFailed, "LIST"))
	}

	defer conn.Close()

	s.handleResponse(AcceptedDataConnection)

	switch a := len(arg); a {
	// list current working directory
	case 1:
		cwd, _ := os.Getwd()
		files, err := ioutil.ReadDir(cwd)
		if err != nil {
			s.handleResponse(RequestedFileActionNotTaken)
		}

		for _, file := range files {
			fmt.Fprintf(conn, "%s\r\n", file.Name())
		}
	}

	s.handleResponse(ClosingDataConnection)
}
