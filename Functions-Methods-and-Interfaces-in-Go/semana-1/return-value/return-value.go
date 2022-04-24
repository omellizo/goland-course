package main

import "fmt"

func foo(x int) int {
	return x + 1
}

func main() {
	y := foo(8)
	fmt.Print(y)
}
