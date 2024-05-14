package main

import (
	"fmt"
	"time"
)

func showMsg(msg string) {
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(msg)
	}
}

func main() {
	go showMsg("Go is a great programming language")
	showMsg("Welcome! Gophers")
}
