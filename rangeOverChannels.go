package main

import "fmt"

func rangeOverChannelsTest() {

	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"

	close(queue)

	// even if we close the channel, we can still range over it
	for elem := range queue {
		fmt.Println(elem)
	}

}
