package main

import (
	"fmt"
	"sync"
)

var x = 0

func inc(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	x = x + 1
	m.Unlock()
	wg.Done()
}

func main() {
	var w sync.WaitGroup
	var m sync.Mutex

	for i := 0; i < 1000; i++ {
		w.Add(1)
		// mutex needs to be passed as a memory, not as value
		// if it's beeing passed by value, then the race condition will occur again
		go inc(&w, &m)
	}

	w.Wait()
	fmt.Println("final value of x:", x)
}

// without mutex the final value of x changes every run, cause the goroutines access x simuntaneously
// and not in a queue manner
// with mutex the output of x is 1000, everytime!
