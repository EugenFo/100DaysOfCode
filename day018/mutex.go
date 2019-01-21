package main

import (
	"fmt"
	"sync"
	"time"
)

type safeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

func (c *safeCounter) inc(key string) {
	// locks the access of the goroutines
	c.mux.Lock()
	c.v[key]++
	c.mux.Unlock()
}

func (c *safeCounter) val(key string) int {
	c.mux.Lock()
	// locks the access and defers it, so that after the return the mux is again unlocked
	defer c.mux.Unlock()
	return c.v[key]
}

func main() {
	c := safeCounter{v: make(map[string]int)}

	for i := 0; i < 1000; i++ {
		go c.inc("something")
	}

	time.Sleep(time.Second)
	fmt.Println(c.val("something"))
}

// without mux.Lock() and mux.Unlock() the goroutines would access simuntaneously the map and it would run into an error
