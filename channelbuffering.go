package main

import "fmt"

func channelBufferingTest() {
	// if the buffer size is less than the expected recive items
	// we will get a fatal error deadlock
	// by default channels are unbuffered
	// buffered channels accepts a limited number of values with or wihtout a correspoding reciever
	messages := make(chan string, 2)

	messages <- "hello"
	messages <- "world"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
