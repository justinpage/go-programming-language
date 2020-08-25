// Fetch prints the content at a URL
// Functionality was copied from go/src/cmd/go/internal/get
// Command: go run fetch.go http://gopl.io/ch1/helloworld\?go-get\=1
// Answer: Github - https://github.com/adonovan/gopl.io
package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		defer resp.Body.Close()

		imports, err := parseMetaGoImports(resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		for i := range imports {
			fmt.Fprintf(os.Stdout, "Repo: %s\n", imports[i].RepoRoot)
		}
	}
}

type metaImport struct {
	Prefix, VCS, RepoRoot string
}

func parseMetaGoImports(r io.Reader) ([]metaImport, error) {
	d := xml.NewDecoder(r)
	d.CharsetReader = charsetReader
	d.Strict = false
	var imports []metaImport
	for {
		t, err := d.RawToken()
		if err != nil {
			if err != io.EOF && len(imports) == 0 {
				return nil, err
			}
			break
		}
		if e, ok := t.(xml.StartElement); ok && strings.EqualFold(e.Name.Local, "body") {
			break
		}
		if e, ok := t.(xml.EndElement); ok && strings.EqualFold(e.Name.Local, "head") {
			break
		}
		e, ok := t.(xml.StartElement)
		if !ok || !strings.EqualFold(e.Name.Local, "meta") {
			continue
		}
		if attrValue(e.Attr, "name") != "go-import" {
			continue
		}
		if f := strings.Fields(attrValue(e.Attr, "content")); len(f) == 3 {
			imports = append(imports, metaImport{
				Prefix:   f[0],
				VCS:      f[1],
				RepoRoot: f[2],
			})
		}
	}

	// Note: ignoring go mod support; focusing on GOPATH for exercise only
	return imports, nil
}

// charsetReader returns a reader that converts from the given charset to UTF-8.
// Currently it only supports UTF-8 and ASCII. Otherwise, it returns a meaningful
// error which is printed by go get, so the user can find why the package
// wasn't downloaded if the encoding is not supported. Note that, in
// order to reduce potential errors, ASCII is treated as UTF-8 (i.e. characters
// greater than 0x7f are not rejected).
func charsetReader(charset string, input io.Reader) (io.Reader, error) {
	switch strings.ToLower(charset) {
	case "utf-8", "ascii":
		return input, nil
	default:
		return nil, fmt.Errorf("can't decode XML document using charset %q", charset)
	}
}

// attrValue returns the attribute value for the case-insensitive key
// `name', or the empty string if nothing is found.
func attrValue(attrs []xml.Attr, name string) string {
	for _, a := range attrs {
		if strings.EqualFold(a.Name.Local, name) {
			return a.Value
		}
	}
	return ""
}
