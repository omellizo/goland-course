package main

import "fmt"

func incFn(x int) int {
	return x + 1
}

func main() {
	funcVar := incFn
	fmt.Println(funcVar(1))
}
