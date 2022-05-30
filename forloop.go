package main

import "fmt"

func forloop() {
	i := 1
	for i < 5 {
		fmt.Println(i)
		i++
	}

	fmt.Println("-------------")

	for j := 0; j < 4; j++ {
		fmt.Println(j)
	}

	k := 5
	for {
		fmt.Println("break")
		if k == 8 {
			break
		}
		k++
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

}
