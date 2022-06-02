package main

import "fmt"

type rect struct {
	width  int
	height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r *rect) perimeter() int {
	return 2 * (r.height + r.height)
}

func changeHeight(r rect) { // will make change to the reference object
	r.height = 34
}

func changeHeightByPointer(r *rect) { // will make chnage to the actual object
	r.height = 44
}

func methodstest() {
	rr := rect{5, 6}
	fmt.Println(rr.area())
	fmt.Println(rr.perimeter())
	changeHeight(rr)
	fmt.Println(rr.height)
	changeHeightByPointer(&rr)
	fmt.Println(rr.height)
}
