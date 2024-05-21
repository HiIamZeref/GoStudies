package main

import "fmt"

type triangle struct {
	base float64
	height float64
}

type square struct {
	sideSize float64
}

type shape interface {
	getArea() float64
}

func main() {
	triangle := triangle{base: 3, height: 5}
	square := square{sideSize: 4}

	printArea(triangle)
	printArea(square)

	
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideSize * s.sideSize
}