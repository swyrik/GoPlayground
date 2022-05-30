package main

import (
	"fmt"
	"math"
)

const n = 223222223222222

func constants() {
	const s string = "hello"
	fmt.Println(s)
	// s = "great" can't assign the value twice for a constant

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(math.Sin(1.7))
}
