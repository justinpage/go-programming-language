// Package xkcd provides a Go API for the XKCD comics
// See https://xkcd.com/json.html

package xkcd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type comic struct {
	Number      int    `json:"num"`
	URL         string `json:"img"`
	Title       string `json:"safe_title"`
	Transcript  string `json:"transcript"`
	Alternative string `json:"alt"`
	Error       error  `json:"error"`
}

// SeedIndex indexes all xkcd comics using a worker pool
func SeedIndex() error {
	const indexURL = "https://xkcd.com/info.0.json"
	const concurrency = 50

	resp, err := http.Get(indexURL)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("getting index failed: %s", resp.Status)
	}

	var index struct {
		Number int `json:"num"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&index); err != nil {
		return fmt.Errorf("unmarhaling index failed: %s", err)
	}

	comics := make(chan int)
	results := make(chan *comic, index.Number)

	// Create a max of 200 workers that are ready to process all comics
	for i, j := 1, concurrency; i <= j; i++ {
		go worker(comics, results)
	}

	// Queue all available comics that we need to process from index
	for i, j := 1, index.Number; i <= j; i++ {
		comics <- i
	}

	var library []*comic

	for i, j := 1, index.Number; i <= j; i++ {
		library = append(library, <-results)
	}

	seed, err := json.MarshalIndent(library, "", " ")
	if err != nil {
		return fmt.Errorf("marshaling library failed: %s", err)
	}

	ioutil.WriteFile("/tmp/xkcd.json", seed, 0644)
	if err != nil {
		return fmt.Errorf("creating index failed: %s", err)
	}

	return nil
}

// SeedIndexAlt indexes all xkcd comics without kernel limit
// Read https://stackoverflow.com/q/12952833/2395590 for more details
func SeedIndexAlt() error {
	const IndexURL = "https://xkcd.com/info.0.json"

	resp, err := http.Get(IndexURL)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("getting index failed: %s", resp.Status)
	}

	var index struct {
		Number int `json:"num"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&index); err != nil {
		return fmt.Errorf("unmarhaling index failed: %s", err)
	}

	ch := make(chan *comic)

	for i, j := 1, index.Number; i <= j; i++ {
		go fetchAlt(i, ch)
	}

	var library []*comic

	for i, j := 1, index.Number; i <= j; i++ {
		library = append(library, <-ch)
	}

	seed, err := json.MarshalIndent(library, "", " ")
	if err != nil {
		return fmt.Errorf("marshaling library failed: %s", err)
	}

	ioutil.WriteFile("/tmp/xkcd.json", seed, 0644)
	if err != nil {
		return fmt.Errorf("creating index failed: %s", err)
	}

	return nil
}

func worker(comics <-chan int, results chan<- *comic) {
	for comicId := range comics {
		results <- fetch(comicId)
	}
}

func fetch(comicId int) *comic {
	comicURL := "https://xkcd.com/" + strconv.Itoa(comicId) + "/info.0.json"

	resp, err := http.Get(comicURL)
	if err != nil {
		return &comic{Error: err}
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return &comic{
			Error: fmt.Errorf(
				"getting index failed: %s for comic %v", resp.Status, comicId,
			),
		}
	}

	var webcomic comic
	if err := json.NewDecoder(resp.Body).Decode(&webcomic); err != nil {
		return &comic{
			Error: fmt.Errorf(
				"unmarshaling index failed: %s for comic %v", err, comicId,
			),
		}
	}

	return &webcomic
}

func fetchAlt(comicId int, ch chan<- *comic) {
	comicURL := "https://xkcd.com/" + strconv.Itoa(comicId) + "/info.0.json"

	resp, err := http.Get(comicURL)
	if err != nil {
		ch <- &comic{Error: err}
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		ch <- &comic{
			Error: fmt.Errorf("getting index failed: %s", resp.Status),
		}
		return
	}

	var webcomic comic
	if err := json.NewDecoder(resp.Body).Decode(&webcomic); err != nil {
		ch <- &comic{
			Error: fmt.Errorf("unmarhaling index failed: %s", err),
		}
		return
	}

	ch <- &webcomic
}
