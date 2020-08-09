package memotest

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"testing"
	"time"
)

func httpGetBody(url string, done <-chan struct{}) (interface{}, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request %s: %v", url, err)
	}
	req.Cancel = done

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

var HTTPGetBody = httpGetBody

type request struct {
	url  string
	done chan struct{}
}

func incomingURLs() <-chan request {
	ch := make(chan request)
	go func() {
		for i, req := range []request{
			{"https://golang.org", make(chan struct{})},
			{"https://godoc.org", make(chan struct{})},
			{"https://play.golang.org", make(chan struct{})},
			{"http://gopl.io", make(chan struct{})},
			{"https://golang.org", make(chan struct{})},
			{"https://godoc.org", make(chan struct{})},
			{"https://play.golang.org", make(chan struct{})},
			{"http://gopl.io", make(chan struct{})},
		} {
			if i%3 == 0 {
				close(req.done) // close a few channels, but let others complete
			}
			ch <- req
		}
		close(ch)
	}()
	return ch
}

type M interface {
	Get(key string, done <-chan struct{}) (interface{}, error)
}

func Sequential(t *testing.T, m M) {
	for req := range incomingURLs() {
		start := time.Now()
		value, err := m.Get(req.url, req.done)
		if err != nil {
			t.Log(req.url, err)
			continue
		}
		t.Logf("%s, %s, %d bytes",
			req.url, time.Since(start), len(value.([]byte)))
	}
}

func Concurrent(t *testing.T, m M) {
	var n sync.WaitGroup
	for req := range incomingURLs() {
		n.Add(1)
		go func(req request) {
			defer n.Done()
			start := time.Now()
			value, err := m.Get(req.url, req.done)
			if err != nil {
				t.Log(req.url, err)
				return
			}
			t.Logf("%s, %s, %d bytes",
				req.url, time.Since(start), len(value.([]byte)))
		}(req)
	}
	n.Wait()
}
