package tar

import (
	"archive/tar"
	"fmt"
	"io"
	"log"

	"github.com/justinpage/go-programming-language/10-packages-and-the-go-tool/exercises/10.2.archive/archive"
)

type tape struct {
	*tar.Reader
}

func magic(r io.Reader) bool {
	tr := tar.NewReader(r)
	_, err := tr.Next()
	if err == io.EOF {
		return false
	}
	if err != nil {
		return false
	}
	return true
}

func (r tape) Print(w io.Writer) {
	for {
		hdr, err := r.Next()
		if err == io.EOF {
			break // End of archive
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Contents of %s:\n", hdr.Name)
		if _, err := io.Copy(w, r); err != nil {
			log.Fatal(err)
		}
		fmt.Println()
	}
}

func Decode(r io.Reader) archive.Archive {
	return tape{tar.NewReader(r)}
}

func init() {
	archive.RegisterFormat("tar", magic, Decode)
}
