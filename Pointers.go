package main

import "fmt"

func pointerstest() {

	i := 1
	fmt.Println(i)

	fmt.Println("memory address of i is &i :", &i)

	changeTheValue(&i) // since we are passing the pointer address the values is getting update directly
	fmt.Println(i)

}

func changeTheValue(ptr *int) {
	*ptr = 33
}
