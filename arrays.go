package main

import "fmt"

func arraystest() {
	var a [5]int
	fmt.Println(a)

	a[4] = 100
	fmt.Println(a)

	fmt.Println("get 4th index", a[4])
	fmt.Println("length of the arary is", len(a))

	b := [5]int{1, 2, 4, 3, 5} //declare and initialize the array in one line
	fmt.Println(b)

	var two [3][4]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			two[i][j] = i * j
		}
	}

	fmt.Println(two)

}
