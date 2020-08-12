// The jpeg command reads a PNG image from the standard input
// and writes it as a JPEG image to the standard output.
//
// cat images/pikachu.jpeg | go run decode.go -format=png > test.png
package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"strings"
)

func main() {
	var format string
	flag.StringVar(&format, "format", "jpeg", "output format")
	img, err := Convert(os.Stdin, format)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}
	io.Copy(os.Stdout, img)
}

func Convert(in io.Reader, format string) (io.Reader, error) {
	img, _, err := image.Decode(in)
	if err != nil {
		return nil, err
	}
	format = strings.ToLower(format)
	switch format {
	case "jpeg":
		return toJPEG(img)
	case "jpg":
		return toJPEG(img)
	case "png":
		return toPNG(img)
	case "gif":
		return toGIF(img)
	default:
		return nil, errors.New("image: unknown format " + format)
	}
}

func toJPEG(img image.Image) (io.Reader, error) {
	buf := new(bytes.Buffer)
	if err := jpeg.Encode(buf, img, &jpeg.Options{Quality: 95}); err != nil {
		return nil, err
	}
	return buf, nil
}

func toPNG(img image.Image) (io.Reader, error) {
	buf := new(bytes.Buffer)
	if err := png.Encode(buf, img); err != nil {
		return nil, err
	}
	return buf, nil
}

func toGIF(img image.Image) (io.Reader, error) {
	buf := new(bytes.Buffer)
	if err := gif.Encode(buf, img, &gif.Options{}); err != nil {
		return nil, err
	}
	return buf, nil
}
