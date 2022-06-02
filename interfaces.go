package main

import (
	"fmt"
	"math"
)

func interfacestest() {
	measure(&square{4})
	measure(&circle{22})
}

// a struct must implement all the methods if it needs to follow the interface
type gemotery interface {
	area() int
	perimeter() int
}

type square struct {
	side int
}

type circle struct {
	radius int
}

func (s *square) area() int {
	return s.side * s.side
}

func (s *square) perimeter() int {
	return s.side * 4
}

func (c *circle) area() int {
	return 2 * c.radius * int(math.Trunc(math.Pi))
}

func (c *circle) perimeter() int {
	return int(math.Trunc(math.Pi)) * c.radius * c.radius
}

func measure(g gemotery) {
	fmt.Println(g.area())
	fmt.Println(g.perimeter())
}
