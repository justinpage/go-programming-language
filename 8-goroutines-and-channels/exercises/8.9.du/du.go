package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var verbose = flag.Bool("v", false, "show verbose progress messages")
var sema = make(chan struct{}, 20)

type node struct {
	path           string
	nfiles, nbytes int64
	wg             *sync.WaitGroup
}

type tree []*node

func main() {
	// Determine the initial directories
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		args = []string{"."}
	}

	// Traverse each root of the file tree in parallel
	var wg sync.WaitGroup
	var roots tree
	fileSizes := make(chan map[string]int64)
	for _, v := range args {
		n := &node{path: v, wg: &wg}
		roots = append(roots, n)
		n.wg.Add(1)
		go walkDir(v, n, fileSizes)
	}

	go func() {
		wg.Wait()
		close(fileSizes)
	}()

	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}
loop:
	for {
		select {
		case size, ok := <-fileSizes:
			if !ok {
				break loop
			}
			for k, v := range size {
				for i, r := range roots {
					if r.path == k {
						roots[i].nfiles++
						roots[i].nbytes += v
					}
				}
			}
		case <-tick:
			printDiskUsage(roots)
		}
	}
	printDiskUsage(roots)
}

func printDiskUsage(r tree) {
	for _, v := range r {
		fmt.Printf("%s:\t%d files\t%.1f GB\n",
			v.path, v.nfiles, float64(v.nbytes)/1e9)
	}
}

// walkDir recursively walks the file tree rooted at dir
// and sends the size of each found file on fileSizes
func walkDir(root string, n *node, fs chan<- map[string]int64) {
	defer n.wg.Done()
	for _, entry := range dirents(n.path) {
		if entry.IsDir() {
			subdir := filepath.Join(n.path, entry.Name())
			n := &node{path: subdir, wg: n.wg}
			n.wg.Add(1)
			go walkDir(root, n, fs)
		} else {
			fs <- map[string]int64{root: entry.Size()}
		}
	}
}

// dirents returns the entries of directory p.
func dirents(p string) []os.FileInfo {
	sema <- struct{}{}        // acquire token
	defer func() { <-sema }() // release token
	entries, err := ioutil.ReadDir(p)
	if err != nil {
		fmt.Fprintln(os.Stderr, "du1: %v\n", err)
	}
	return entries
}
