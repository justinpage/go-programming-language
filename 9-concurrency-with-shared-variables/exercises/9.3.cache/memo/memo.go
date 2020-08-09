// Package memo provides a concurrency-unsafe
// memoization of a function of type Func
package memo

import (
	"log"
	"sync"
)

// A Memo caches the results of calling a Func.
type Memo struct {
	f          Func
	cache      map[string]*entry
	sync.Mutex // guards cache
}

// Func is the type of the function to memoize.
type Func func(key string, done <-chan struct{}) (interface{}, error)

type entry struct {
	res   result
	ready chan struct{}   // closed when res is ready
	done  <-chan struct{} // support cancellation
}

type result struct {
	value interface{}
	err   error
}

func NewMemo(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]*entry)}
}

func (memo *Memo) Get(key string, done <-chan struct{}) (value interface{}, err error) {
	memo.Lock()
	e := memo.cache[key]
	if e == nil {
		// This is the first request for this key.
		// This goroutine becomes responsible for computing
		// the value and broadcasting the ready condition
		e = &entry{ready: make(chan struct{}), done: done}
		memo.cache[key] = e
		memo.Unlock()

		e.res.value, e.res.err = memo.f(key, done)
		if e.res.err != nil {
			memo.Lock()
			delete(memo.cache, key)
			memo.Unlock()
			return nil, e.res.err
		}

		close(e.ready) // broadcast ready condition
	} else {
		// This is a repeat request for this key.
		memo.Unlock()
		// Receiving on "e.done" from an entry makes it possible for us to retry
		// a request that wouldn't complete normally because it is waiting for
		// another job to finish --But that job was unexpectedly canceled.
		// Notice how we are receiving on "e.done" while still allowing
		// ourselves to use the done passed through the original memo.Get call.
		//
		// In other words,"e.done" is in reference to an already existing value
		// from the map (meaning another job is handling it but sometimes is
		// canceled) while "done" is in reference to the done channel passed for
		// that specific call
		for {
			select {
			case <-e.done:
				log.Println("retrying...")
				return memo.Get(key, done)
			case <-e.ready:
				log.Println("getting existing")
				return e.res.value, e.res.err
			}
		}
	}
	return e.res.value, e.res.err
}
