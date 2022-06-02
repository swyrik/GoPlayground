package main

import (
	"fmt"
	"sync"
	"time"
)

func waitgroupstest() {

	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {

		j := i // if we don't bind this to new varibale we will see duplicates
		/*
			by the time the nth o routine has started the i variable value might have been changed
			so we might observe duplicates here. just like using var with setTimeOut in javascript
		*/

		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(j)
		}()
	}

	wg.Wait()

}

func worker(id int) {
	fmt.Printf("worker %d starting\n", id)
	time.Sleep(10)
	fmt.Printf("worker %d done\n", id)
}
