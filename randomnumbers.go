package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomNumsTest() {
	// the below code will result in same set of results everytime
	fmt.Println(rand.Intn(100), ",", rand.Intn(100))
	fmt.Println(rand.Float64())
	fmt.Println((rand.Float64()*5)+5, ",", (rand.Float64()*5)+5)

	// havina a new seed everytime will result int random values as output
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	fmt.Printf("===================\n\n")
	fmt.Println(r1.Intn(100), ",", r1.Intn(100))
	fmt.Println(r1.Float64())
	fmt.Println((r1.Float64()*5)+5, ",", (r1.Float64()*5)+5)

	fmt.Printf("===================\n\n")
	// having same source will result in same random numbers
	s2 := rand.NewSource(42)
	r2 := rand.New(s2)
	fmt.Print(r2.Intn(100), ",")
	fmt.Print(r2.Intn(100))
	fmt.Println()
	s3 := rand.NewSource(42)
	r3 := rand.New(s3)
	fmt.Print(r3.Intn(100), ",")
	fmt.Print(r3.Intn(100))
}
