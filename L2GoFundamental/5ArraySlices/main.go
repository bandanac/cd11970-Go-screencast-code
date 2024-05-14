package main

import "fmt"

func main() {
	fiboA := [8]int{0, 1, 1, 2, 3, 5, 8, 13}
	fmt.Println(fiboA)

	fiboS := []int{0, 1, 1, 2, 3, 5, 8, 13}
	fmt.Println(fiboS)

	fmt.Println(fiboS[0])
	fmt.Println(fiboS[7])
	fmt.Println(len(fiboS))
	fmt.Println(fiboS[0:4])

	fiboS = append(fiboS, 21)

	fmt.Println(fiboS)
}
