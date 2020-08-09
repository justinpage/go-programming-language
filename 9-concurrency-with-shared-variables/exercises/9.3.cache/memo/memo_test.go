package memo

import (
	"testing"

	"github.com/justinpage/go-programming-language/9-concurrency-with-shared-variables/exercises/9.3.cache/memotest"
)

var httpGetBody = memotest.HTTPGetBody

func Test(t *testing.T) {
	m := NewMemo(httpGetBody)
	memotest.Sequential(t, m)
}

func TestConcurrent(t *testing.T) {
	m := NewMemo(httpGetBody)
	memotest.Concurrent(t, m)
}
