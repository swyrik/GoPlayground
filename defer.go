package main

import (
	"fmt"
	"os"
)

func defertest() {

	f := createFile("C:/Users/GRAVITY/Desktop/sample.txt")
	defer closeFile(f) // this function will execute in the last of main stack call
	writeFile(f)

}

func createFile(p string) *os.File {
	fmt.Println("creating file")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing to file")
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("closing file")
	err := f.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
