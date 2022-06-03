package main

import (
	"fmt"
	"time"
)

func channelsTimeOutTest() {

	c1 := make(chan string, 1)
	c2 := make(chan string, 1)

	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result1"
	}()

	select {
	case msg1 := <-c1:
		fmt.Println(msg1)
	case <-time.After(time.Second):
		fmt.Println("timeout 1")
	}

	go func() {
		time.Sleep(3 * time.Second)
		c2 <- "result2"
	}()

	select {
	case msg2 := <-c2:
		fmt.Println(msg2)
	case <-time.After(4 * time.Second):
		fmt.Println("timeout 2")
	}
}
