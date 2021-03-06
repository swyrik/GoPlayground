package main

import (
	"fmt"
	"time"
)

func rateLimiter() {

	requests := make(chan int, 5)
	for i := 0; i < 5; i++ {
		requests <- i
	}

	close(requests)

	limiter := time.Tick(1000 * time.Millisecond)

	// we are alowing one request for every one second
	// by blocking with time ticker "limiter"
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	burstyLimiter := make(chan time.Time, 3)
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(1000 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	fmt.Println("")
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
