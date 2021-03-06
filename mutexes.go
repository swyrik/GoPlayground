package main

import (
	"fmt"
	"sync"
)

type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

// if we don't use mutex we will get concurrent write operations erroor
func (c *Container) inc(name string) {
	c.mu.Lock()
	// this defer will be performed after the counter has incemented
	defer c.mu.Unlock()
	c.counters[name]++
}

func mutexttest() {
	c := Container{
		counters: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup

	doIncrement := func(name string, n int) {
		for i := 0; i < n; i++ {
			c.inc(name)
		}
		wg.Done()
	}

	wg.Add(3)
	go doIncrement("a", 10000)
	go doIncrement("a", 10000)
	go doIncrement("b", 10000)

	wg.Wait()
	fmt.Println(c.counters)

}
