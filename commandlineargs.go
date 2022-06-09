package main

import (
	"fmt"
	"os"
)

func commandlineargstest() {

	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	arg := os.Args[3]

	fmt.Println(argsWithProg, "\n", argsWithoutProg, "\n", arg, "\n", os.Args[0])

}
