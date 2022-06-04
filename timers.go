package main

import (
	"fmt"
	"time"
)

func timers() {

	timer1 := time.NewTimer(2 * time.Second)

	currentTime := <-timer1.C

	fmt.Println("timer 1 fired", currentTime)

	timer2 := time.NewTimer(3 * time.Second)
	go func() {
		timenow := <-timer2.C
		fmt.Println("timer 2 fired", timenow)
	}()

	stop := timer2.Stop()
	if stop {
		fmt.Println("timer 2 stopped")
	}
}
