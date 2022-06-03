package main

import (
	"fmt"
	"time"
)

func channelsynctest() {
	done := make(chan bool, 1)
	go workerChan(done)

	// if we remove the below line
	// program will immediatly exit without
	// starting the workerChan method
	<-done
}

func workerChan(done chan bool) {
	fmt.Println("started working")
	time.Sleep(time.Second)
	fmt.Println("work completed")
	done <- true
}
