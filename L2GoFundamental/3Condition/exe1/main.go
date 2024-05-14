package main

import "fmt"

func main() {
	number := 100

	if number < 0 {
		fmt.Println(number, "is negative")
	} else if number < 100 && number > 0 {
		fmt.Println(number, "is positive")
	} else {
		fmt.Println(number, "is positive and is a large number!")
	}
}
