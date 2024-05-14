package main

import (
	"fmt"
	"strconv"
)

type Car struct {
	make string
	year int
	used bool
}

// This method is a value receiver
func (c Car) describe() string {
	used := ""

	if c.used {
		used = "a used car"
	} else {
		used = "a new car"
	}

	return "This " + strconv.Itoa(int(c.year)) + " " + c.make + " is " + used
}

func main() {
	car1 := Car{make: "DoLerean", year: 1985, used: true}
	car2 := Car{make: "Honda", year: 2023, used: false}

	fmt.Println(car1)
	fmt.Println(car2)

	fmt.Println(car1.describe())
	fmt.Println(car1.year)
	fmt.Println(car2.describe())
	fmt.Println(car2.make)
}
