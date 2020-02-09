// References:

// https://cr.yp.to/ftp.html
// https://github.com/torbiak/gopl/tree/master/ex8.2
// https://github.com/torbiak/gopl/tree/master/ex8.2
// https://github.com/kspviswa/lsgo/blob/master/ls.go
package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/user"
	"strconv"
	"strings"
	"syscall"
	"text/tabwriter"
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
	UserLoggedInProceed         = "230 User logged in, proceed\n"
	PathNameCreated             = "257 Created \"%s\"\n"
	CurrentWorkingDirectory     = "257 \"%s\"\n"
	UserOkayNeedPassword        = "331 User %s okay, need password\n"
	RequestedFileActionNotTaken = "450 Requested file action not taken\n"
	RequestedActionHasFailed    = "500 Requested action has failed \"%s\"\n"
	CommandNotImplemented       = "502 Command not implemented \"%s\"\n"
)

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

	temp, err := ioutil.TempDir("/tmp/", "ftp-")
	if err != nil {
		return // unable to create temporary directory
	}

	seedFolder(temp)
	if err != nil {
		return // unable to use temporary directory
	}

	defer os.RemoveAll(temp)

	cmd := bufio.NewScanner(s.conn)
	for cmd.Scan() {
		cmd := cmd.Text()
		arg := strings.Split(cmd, " ")

		if len(arg) > 1 {
			cmd = arg[0]
		}

		switch cmd {
		case "USER":
			s.handleResponse(fmt.Sprintf(UserOkayNeedPassword, arg[0]))
		case "PASS":
			s.handleResponse(UserLoggedInProceed)
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
		case "NLST":
			s.handleNLST(arg)
		case "PWD":
			cwd, _ := os.Getwd()
			s.handleResponse(fmt.Sprintf(CurrentWorkingDirectory, cwd))
		default:
			fmt.Println("cmd", cmd)
			s.handleResponse(fmt.Sprintf(CommandNotImplemented, cmd))
		}
	}

	if cmd.Err() != nil {
		return // something went wrong (not io.EOF); ignore for now
	}
}

func seedFolder(temp string) error {
	dat := []byte("hello\nftp\n")
	err := ioutil.WriteFile(temp+"/message.md", dat, 0666)
	if err != nil {
		return err
	}

	err = os.Mkdir(temp+"/server", 0755)
	if err != nil {
		return err
	}

	dat, err = ioutil.ReadFile("./server.go")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(temp+"/server/main.go", dat, 0666)
	if err != nil {
		return err
	}

	err = os.Chdir(temp) // start each connection inside /tmp/ dir
	if err != nil {
		return err
	}

	return nil
}

type server struct {
	conn net.Conn
	pasv net.Listener
	text *tabwriter.Writer
}

func (s *server) handleResponse(msg string) {
	_, err := io.WriteString(s.conn, msg)
	if err != nil {
		return // e.g., client disconnected
	}
}

func (s *server) handlePASV() {
	// NEED: better error handling
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

	tw := new(tabwriter.Writer).Init(conn, 0, 8, 2, ' ', 0)
	list := func(file os.FileInfo) {
		const format = "%s\t%3v %s\t%s\t%12v %s %s\r\n"

		mode := file.Mode().String()
		link := file.Sys().(*syscall.Stat_t).Nlink

		uid := strconv.Itoa(int(file.Sys().(*syscall.Stat_t).Uid))
		owner, _ := user.LookupId(uid)
		username := owner.Username

		gid := strconv.Itoa(int(file.Sys().(*syscall.Stat_t).Gid))
		group, _ := user.LookupGroupId(gid)
		groupname := group.Name

		size := file.Size()
		time := file.ModTime().Format("Jan  2 15:04")
		name := file.Name()

		fmt.Fprintf(tw, format, mode, link, username,
			groupname, size, time, name,
		)

		tw.Flush()
	}

	switch a := len(arg); a {
	// list current working directory
	case 1:
		cwd, _ := os.Getwd()
		files, err := ioutil.ReadDir(cwd)
		if err != nil {
			s.handleResponse(RequestedFileActionNotTaken)
		}

		for _, file := range files {
			list(file)
		}
	// list specific file or directory
	case 2:
		cwd, _ := os.Getwd()
		files, err := ioutil.ReadDir(cwd)
		if err != nil {
			s.handleResponse(RequestedFileActionNotTaken)
		}

		match := arg[1]
		for _, file := range files {
			if file.Name() == match {
				list(file)
			}
		}
	}

	s.handleResponse(ClosingDataConnection)
}

func (s *server) handleNLST(arg []string) {
	conn, err := s.pasv.Accept()
	if err != nil {
		log.Print(err) // e.g., connection aborted
		s.handleResponse(fmt.Sprintf(RequestedActionHasFailed, "NLST"))
	}

	defer conn.Close()

	s.handleResponse(AcceptedDataConnection)

	cwd, _ := os.Getwd()
	files, err := ioutil.ReadDir(cwd)
	if err != nil {
		s.handleResponse(RequestedFileActionNotTaken)
	}

	for _, file := range files {
		fmt.Fprintf(conn, "%s\r\n", file.Name())
	}

	s.handleResponse(ClosingDataConnection)
}
