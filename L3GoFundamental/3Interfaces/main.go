package main

import (
	"fmt"
)

// A "Baseball" struct with two fields
type Baseball struct {
	mass         int
	acceleration int
}

// A "Football" struct with no fields
type Football struct{}

// A "getForce" method that accepts a "Baseball" type
func (b Baseball) getForce() int {
	return b.mass * b.acceleration
}

// A "getForce" method that accepts a "Football" type
func (f Football) getForce() int {
	return 50
}

type Force interface {
	getForce() int
}

func compareForce(a, b Force) bool {
	return a.getForce() > b.getForce()
}

func main() {

	b := Baseball{
		mass:         2,
		acceleration: 20,
	}

	f := Football{}

	fmt.Println(compareForce(b, f))

}
