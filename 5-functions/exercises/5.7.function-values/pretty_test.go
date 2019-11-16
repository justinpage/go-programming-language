package main

import (
	"testing"

	"golang.org/x/net/html"
)

func TestOutlineParse(t *testing.T) {
	output := Outline("http://gopl.io")
	_, err := html.Parse(&output)
	if err != nil {
		t.Errorf("Outline(htto://gopl.io) failed")
	}
}
