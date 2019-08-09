package main

import (
	"fmt"
)

type area interface {
	getArea() float64
}

type triangle struct {
	base   float64
	height float64
}
type square struct {
	side float64
}

func main() {
	t := triangle{
		base:   12.6,
		height: 25.7,
	}
	s := square{
		side: 9.9,
	}

	printArea(t)
	printArea(s)
}

func printArea(a area) {
	fmt.Println(a.getArea())
}

func (t triangle) getArea() float64 {
	// VER CUSTOM LOGIC FOR GENERATING AN ENGLISH GREETING
	return (0.5 * t.base * t.height)
}

func (s square) getArea() float64 {
	return (s.side * s.side)
}
