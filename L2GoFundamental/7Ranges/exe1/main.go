package main

import (
	"fmt"
)

func reduce(list []int) int {
	sum := 0

	for _, temp := range list {
		sum += temp
	}

	return sum
}

func main() {
	fmt.Println(reduce(
		[]int{0, 1, 1, 2, 3, 5, 8, 13, 21}))
}
