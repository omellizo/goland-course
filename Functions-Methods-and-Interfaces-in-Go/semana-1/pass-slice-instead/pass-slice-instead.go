package main

import "fmt"

func foo(slice []int) int {
	slice[0] = slice[0] + 1
	return slice[0]
}

func main() {
	a := []int{1, 2, 3}
	foo(a)
	fmt.Println(a)
}
