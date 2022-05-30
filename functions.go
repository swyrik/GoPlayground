package main

import "fmt"

func functest() {
	fmt.Println(plus(2, 3))
	fmt.Println(plusPlus(2, 4, 8))

	fmt.Println(returnTwo())

	_, b := returnTwo()
	fmt.Println("the second return value from returnTwo func", b)

	_, c, d := returnThree()
	fmt.Printf("the second and third return value from returnThree func '%v' '%v'\n", c, d)
}

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

// functions can return multiple values at a time
func returnTwo() (int, int) {
	return 1, 2
}

func returnThree() (int, string, bool) {
	return 1, "two", true
}
