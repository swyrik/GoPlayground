package main

import (
	"fmt"
	"time"
)

func myPanic() {
	panic("a problem")
}

func recovertest() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered Error:\n", r)
		}
	}()

	myPanic()
	time.Sleep(time.Second)
	fmt.Println("After mayPanic()")

}
