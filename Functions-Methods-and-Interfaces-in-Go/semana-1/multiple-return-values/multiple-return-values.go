package main

import "fmt"

func foo(x int) (int, int, int, int) {
	return x, x + 1, x + 2, x + 3
}

func main() {
	a, b, c, d := foo(8)
	fmt.Print(a, b, c, d)
}