package main

import (
	"fmt"
)

type Shape interface {
	perimeter() int
}

// A "Rectangle" struct with two fields
type Rectangle struct {
	length  int
	breadth int
}

// A "Square" struct with a fields
type Square struct {
	side int
}

func (r Rectangle) perimeter() int {
	return (2 * r.length) + (2 * r.breadth)
}

func (s Square) perimeter() int {
	return 4 * s.side
}

func getPerimeter(s Shape) int {
	return s.perimeter()
}

func main() {
	rectangle := Rectangle{
		length:  2,
		breadth: 4,
	}

	square := Square{
		side: 2,
	}

	fmt.Println(getPerimeter(rectangle))
	fmt.Println(getPerimeter(square))

}
