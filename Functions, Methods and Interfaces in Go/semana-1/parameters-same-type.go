package main

import "fmt"

func foo(x, y int) {
	fmt.Print(x * y)
}

func main() {
	foo(2, 3)
}
