package main

import "fmt"

func foo() *int {
	x := 1
	fmt.Println("x", x)
	fmt.Println("&x", &x)
	return &x
}

func main() {
	var y *int
	y = foo()
	fmt.Println("&y", &y)
	fmt.Println("*y", *y)
	fmt.Println("y", y)
}
