// The command reads an archive from the standard input
// and writes its contents to standard output.
//
// cat testdata/message.tar | go run extract.go
package main

import (
	"fmt"
	"io"
	"os"

	"github.com/justinpage/go-programming-language/10-packages-and-the-go-tool/exercises/10.2.archive/archive"
	_ "github.com/justinpage/go-programming-language/10-packages-and-the-go-tool/exercises/10.2.archive/archive/tar"
	_ "github.com/justinpage/go-programming-language/10-packages-and-the-go-tool/exercises/10.2.archive/archive/zip"
)

func main() {
	arc, err := Extract(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}
	arc.Print(os.Stdout)
}

func Extract(in io.Reader) (archive.Archive, error) {
	arc, _, err := archive.Decode(in)
	if err != nil {
		return nil, err
	}
	return arc, nil
}
