package main

import (
	"fmt"
	"os"
)

func exittest() {

	defer fmt.Println("!") // this will never be printed, since we are exiting the program
	fmt.Println("hello world")
	os.Exit(1) // it exits the program with ret

}
