package main

import "fmt"

func main() {
	i := 1
	defer fmt.Println(i + 1)
	i++
	fmt.Println("Hello")
}
