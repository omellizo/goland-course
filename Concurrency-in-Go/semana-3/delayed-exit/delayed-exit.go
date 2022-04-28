package main

import (
	"fmt"
	"time"
)

func main() {
	go fmt.Println("New Goroutine")
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Main Goroutine")
}