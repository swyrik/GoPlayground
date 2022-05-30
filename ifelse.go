package main

import "fmt"

func ifelse() {

	if 7%2 == 1 {
		fmt.Println("odd number")
	}

	if 10%2 == 0 {
		fmt.Println("even number")
	}

	if num := 9; num < 0 { // variables or const declared in if/else block will be available in the block
		fmt.Println(num, "is greater than zero")
	} else if num < 10 {
		fmt.Println(num, "is less than ten")
	} else {
		fmt.Println("else block")
	}

	// fmt.Print(num) num isn't avilable here which we have declared in if/else block

	// if true
	// 	fmt.Println("true")

	// curly braces are required for sure in ifelse block

	// there is no terneary operator in go lang
}
