package main

import (
	"fmt"
	"time"
)

func tickers() {
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	time.Sleep(3 * time.Second)
	ticker.Stop()
	done <- true
	fmt.Println("ticker stopped")

}
