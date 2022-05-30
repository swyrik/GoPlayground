package main

import "fmt"

func slicestest() {
	s := make([]string, 3)
	fmt.Println(s)

	s[0] = "Java"
	s[1] = "Kotlin"
	s[2] = "Python"

	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	fmt.Println("length of the slice :", len(s))

	s = append(s, "Go")
	s = append(s, "Rust", "JavaScript", "Scala", "TypeScript")

	fmt.Println(s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("Elements from copied array")
	fmt.Println(c)

	l := s[2:5]
	fmt.Println("Elements copied using a range s[2:5]", l)

	l = s[:5] // from first element to 5th element in the slice
	fmt.Println("sl2:", l)

	l = s[2:] // from second element to last element in the slice
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = j * i
		}

	}

	fmt.Println(twoD)

}
