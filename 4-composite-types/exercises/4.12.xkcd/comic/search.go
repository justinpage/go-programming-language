package xkcd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type SearchItem struct {
	Number      int    `json:"num"`
	URL         string `json:"img"`
	Title       string `json:"safe_title"`
	Transcript  string `json:"transcript"`
	Alternative string `json:"alt"`
}

type Result struct {
	Number     int
	URL        string
	Title      string
	Transcript string
}

func Search(term string) ([]*Result, error) {
	xkcd, err := ioutil.ReadFile("/tmp/xkcd.json")
	if err != nil {
		return nil, fmt.Errorf("unmarhaling index failed: %s", err)
	}

	var library []*SearchItem
	if err := json.Unmarshal(xkcd, &library); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}

	var results []*Result
	for _, item := range library {
		switch {
		case strings.Contains(item.Transcript, term):
			match := &Result{item.Number, item.URL, item.Title, item.Transcript}
			results = append(results, match)
		case strings.Contains(item.Alternative, term):
			match := &Result{item.Number, item.URL, item.Title, item.Alternative}
			results = append(results, match)
		case strings.Contains(item.Title, term):
			match := &Result{item.Number, item.URL, item.Title, item.Title}
			results = append(results, match)
		}
	}

	return results, nil
}
