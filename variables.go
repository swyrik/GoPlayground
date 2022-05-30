package main

import "fmt"

func variables() {
	var a = "inital"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e) // zero if not initialized

	var s string
	fmt.Println(s) // empty string if not initialized

	f := "fox" // automatic declaration of variable
	fmt.Println(f)
}
