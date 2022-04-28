package main

import "fmt"

func main() {
	go fmt.Println("New Goroutine")
	fmt.Println("Main Goroutine")
}