package main

import "fmt"

func values() {
	fmt.Println("hello", "world")
	fmt.Println("1+1", 1+1)
	fmt.Println(3 / 4)     // quotient
	fmt.Println(3.0 / 4.0) // fractional value
	fmt.Println(4.0 / 3.0) // fractional value
	fmt.Println(5 % 2)     // remainder
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
