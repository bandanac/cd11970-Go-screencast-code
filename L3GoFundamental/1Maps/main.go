package main

import (
	"fmt"
)

func main() {
	dictionary := map[string]string{
		"Go":     "Programming language created by Google",
		"Gopher": "A software engineer who builds with Go",
		"Golang": "Another name for Go",
	}

	fmt.Println(dictionary)
	fmt.Println(dictionary["Gopher"])

	dictionary["Gopher"] = "The fuzzy mascot for Go"
	dictionary["Map"] = "An unordered data structure with key-value pair"

	fmt.Println(dictionary)

	for word, definition := range dictionary {
		fmt.Println("The definition of", word, "is:", definition)
	}

	delete(dictionary, "Map")

	fmt.Println(dictionary)

}
