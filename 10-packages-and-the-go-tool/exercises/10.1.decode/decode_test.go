package main

import (
	"flag"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	flag.Parse()
	exitCode := m.Run()

	os.Exit(exitCode)
}

func TestConvertPNGtoJPEG(t *testing.T) {
	var path string
	flag.StringVar(&path, "img1", "images/charmander.png", "path to png image")

	png, err := os.Open(path)
	if err != nil {
		t.Fatalf("Could not open image from path: %s got %v", path, err)
	}

	img, err := Convert(png, "jpeg")
	if err != nil {
		t.Fatalf("Could not convert png to jpeg: %v", err)
	}

	if _, err := jpeg.Decode(img); err != nil {
		t.Fatalf("Not a valid jpeg image: %v", err)
	}
}

func TestConvertJPEGtoPNG(t *testing.T) {
	var path string
	flag.StringVar(&path, "img2", "images/pikachu.jpeg", "path to jpeg image")

	jpeg, err := os.Open(path)
	if err != nil {
		t.Fatalf("Could not open image from path: %s got %v", path, err)
	}

	img, err := Convert(jpeg, "png")
	if err != nil {
		t.Fatalf("Could not convert jpeg to png: %v", err)
	}

	if _, err := png.Decode(img); err != nil {
		t.Fatalf("Not a valid png image: %v", err)
	}
}

func TestConvertPNGToGIF(t *testing.T) {
	var path string
	flag.StringVar(&path, "img3", "images/charmander.png", "path to png image")

	png, err := os.Open(path)
	if err != nil {
		t.Fatalf("Could not open image from path: %s got %v", path, err)
	}

	img, err := Convert(png, "gif")
	if err != nil {
		t.Fatalf("Could not convert png to jpeg: %v", err)
	}

	if _, err := gif.Decode(img); err != nil {
		t.Fatalf("Not a valid gif image: %v", err)
	}
}

func TestConvertWEBPToPNG(t *testing.T) {
	var path string
	flag.StringVar(&path, "img4", "images/squirtle.webp", "path to webp image")

	webp, err := os.Open(path)
	if err != nil {
		t.Fatalf("Could not open image from path: %s got %v", path, err)
	}

	_, err = Convert(webp, "png")
	if err == nil {
		t.Fatalf("webp images not support, expected error")
	}
	if err.Error() != "image: unknown format" {
		t.Fatalf("webp images not support, expected error unknown format")
	}
}

func TestConvertJPEGToWEBP(t *testing.T) {
	var path string
	flag.StringVar(&path, "img5", "images/pikachu.jpeg", "path to jpeg image")

	jpeg, err := os.Open(path)
	if err != nil {
		t.Fatalf("Could not open image from path: %s got %v", path, err)
	}

	_, err = Convert(jpeg, "webp")
	if err == nil {
		t.Fatalf("webp images not support, expected error")
	}
	if err.Error() != "image: unknown format webp" {
		t.Fatalf("webp images not support, expected error unknown format")
	}
}
