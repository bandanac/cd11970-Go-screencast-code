package main

import (
	"fmt"
	"strconv"
)

func fizzbuzz(x int) []string {
	result := []string{}

	for i := 1; i <= x; i++ {
		if i%15 == 0 {
			result = append(result, "FizzBuzz")
		} else if i%3 == 0 {
			result = append(result, "Fizz")
		} else if i%5 == 0 {
			result = append(result, "Buzz")
		} else {
			result = append(result, strconv.Itoa(i))
		}
	}

	return result

}
func main() {
	fmt.Println(fizzbuzz(15))
}
