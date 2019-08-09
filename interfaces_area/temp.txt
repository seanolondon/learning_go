package main

import "fmt"

type shape interface {
	printArea
}

type square struct {
	sideLength float64
}

type triangle struct {
	height float64
	base   float64
}

func main() {
	triangle := shape{base: 5.5,
		height: 8.2}
	square := shape{sideLength: 10.1}

	jim.updateName("jimmy")
	jim.print()
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

// func (pointerToPerson *person) updateName(newFirstName string) {
// 	(*pointerToPerson).firstName = newFirstName
// }

func (pointerToTriangle *triangle) getArea(area float64) {
	(0.5 * pointerToTriangle.base * *pointerToTriangle.height) = area
}

func (pointerToSquare *square) getArea(area float64) {
	(*pointerToSquare.sideLength * *pointerToSquare.sideLength) = area
}
