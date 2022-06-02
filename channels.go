package main

import "fmt"

/*
channels are pipes that can connect the goroutines
a recieve end of a channle is a blocking statement
*/

func channelstest() {

	message := make(chan string)

	fmt.Printf("%T\n", message)
	go func() { message <- "ping" }()

	msg := <-message // channel receving is a blocking statement
	fmt.Println(msg)

}
