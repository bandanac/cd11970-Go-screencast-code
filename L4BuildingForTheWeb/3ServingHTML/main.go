package main

import (
	"fmt"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)

	fmt.Println("server is starting at 3000...")
	http.ListenAndServe(":3000", nil)
}
