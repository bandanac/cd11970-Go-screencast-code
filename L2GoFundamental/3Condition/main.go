package main

import "fmt"

func main() {
	language := "Go"

	if language == "Java" {
		fmt.Println("Language is Java")
	} else if language == "C" {
		fmt.Println("Language is C")
	} else {
		fmt.Println("Language is Go")
	}
}
