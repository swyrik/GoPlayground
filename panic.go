package main

import (
	"fmt"
	"os"
)

// panic stops execution of rest of the code

func panictest() {

	panic("a problem")

	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}

	fmt.Println("after panic")

}
