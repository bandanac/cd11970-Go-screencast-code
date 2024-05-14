package main

import (
	"fmt"
)

func main() {
	athletes := []string{"stephen", "kley", "harrison", "draymond", "andrew"}

	for i, name := range athletes {
		fmt.Printf("i = %d, name = %s\n", i, name)
	}

	for _, name := range athletes {
		fmt.Printf("name = %s\n", name)
	}

	numbers := []int{30, 11, 40, 23, 12}

	for _, num := range numbers {
		if num%2 == 0 {
			fmt.Println(num)
		}
	}
}
