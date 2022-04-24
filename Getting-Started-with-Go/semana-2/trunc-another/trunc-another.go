package main

import (
	"fmt"
	"math"
)

func main() {
	var selectedNumber float64
	fmt.Println("Hello! please input a floating number:")
	fmt.Scan(&selectedNumber)
	fmt.Print(math.Trunc(selectedNumber))
}