package main

import "fmt"

func main() {
	var x [5] int
	x[0] = 2
	fmt.Println("x[1] ", x[1])
	fmt.Println("x[1] ", x[0])
}
