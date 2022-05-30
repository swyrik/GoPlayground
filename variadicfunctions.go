package main

import "fmt"

func variadicFunction() {

	fmt.Println(sum(1, 2))
	fmt.Println(sum(3, 4))

	nums := []int{1, 2, 43, 42}
	fmt.Println(sum(nums...)) // this is array spreading

}

// this is similar to var args
func sum(args ...int) int {
	total := 0
	for _, num := range args {
		total += num
	}
	return total
}
