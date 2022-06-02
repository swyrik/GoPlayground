package main

import (
	"fmt"
	"time"
)

func goroutinestest() {
	f("direct")

	go f("concurrent")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")
}

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}
