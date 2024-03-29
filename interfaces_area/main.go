package main

import (
	"fmt"
)

type shape interface {
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

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (t triangle) getArea() float64 {
	//This is different than the chat bot!
	//The chat bot only has something like englishBot
	//because the eb. is not used. Here it is used, so you must
	//put the t.
	return (0.5 * t.base * t.height)
}

func (s square) getArea() float64 {
	return (s.side * s.side)
}
