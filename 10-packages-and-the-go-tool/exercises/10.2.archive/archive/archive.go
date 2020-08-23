package archive

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"sync"
	"sync/atomic"
)

// ErrFormat indicates that decoding encountered an unknown format.
var ErrFormat = errors.New("archive: unknown format")

type Archive interface {
	Print(w io.Writer)
}

// Formats is the list of registered formats.
var (
	formatsMu     sync.Mutex
	atomicFormats atomic.Value
)

// A format holds an arhive format's name, magic header and how to decode it.
type format struct {
	name   string
	magic  func(io.Reader) bool
	decode func(io.Reader) Archive
}

// Sniff determines the format of r's data.
func sniff(r io.Reader) format {
	formats, _ := atomicFormats.Load().([]format)
	for _, f := range formats {
		b := f.magic(r)
		if b {
			return f
		}
	}
	return format{}
}

// Decode decodes an archive that has been encoded in a registered format.
// The string returned is the format name used during format registration.
// Format registration is typically done by an init function in the archive-
// specific package.
func Decode(r io.Reader) (Archive, string, error) {
	b, _ := ioutil.ReadAll(r)
	f := sniff(bytes.NewReader(b))
	if f.decode == nil {
		return nil, "", ErrFormat
	}
	m := f.decode(bytes.NewReader(b))
	return m, f.name, nil
}

// RegisterFormat registers an archive format for use by Decode.
// Name is the name of the format, like "zip" or "tar".
// Magic is the a function that identifies the format's encoding.
// Decode is the function that decodes the encoded archive.
func RegisterFormat(
	name string, magic func(io.Reader) bool,
	decode func(io.Reader) Archive,
) {
	formatsMu.Lock()
	formats, _ := atomicFormats.Load().([]format)
	atomicFormats.Store(append(formats, format{name, magic, decode}))
	formatsMu.Unlock()
}
