package main

import "fmt"

func main() {
	var x [5] int = [5]int{1, 2, 3, 4, 5}
	fmt.Println("X ", x)

	y := [...]int{1, 2, 3, 4}
	fmt.Println("Y ", y)
}
