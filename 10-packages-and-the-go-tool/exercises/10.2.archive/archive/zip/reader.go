package zip

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"log"

	"github.com/justinpage/go-programming-language/10-packages-and-the-go-tool/exercises/10.2.archive/archive"
)

type tape struct {
	*zip.Reader
}

func magic(r io.Reader) bool {
	rr := r.(*bytes.Reader)
	tr, err := zip.NewReader(rr, rr.Size())
	if err != nil {
		return false
	}
	for range tr.File {
		return true
	}
	return false
}

func (r tape) Print(w io.Writer) {
	for _, f := range r.File {
		fmt.Printf("Contents of %s:\n", f.Name)
		rc, err := f.Open()
		if err != nil {
			log.Fatal(err)
		}
		if _, err := io.Copy(w, rc); err != nil {
			log.Fatal(err)
		}
		rc.Close()
		fmt.Println()
	}
}

func Decode(r io.Reader) archive.Archive {
	rr := r.(*bytes.Reader)
	tr, err := zip.NewReader(rr, rr.Size())
	if err != nil {
		return nil
	}
	return tape{tr}
}

func init() {
	archive.RegisterFormat("zip", magic, Decode)
}
