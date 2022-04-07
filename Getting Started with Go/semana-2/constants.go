package main

import "fmt"

func main() {
	const x = 1.3
	const (
		y = 4
		z = "Hi"
	)
	fmt.Println("X ", x)
	fmt.Println("Y ", y)
	fmt.Println("Z ", z)

	type Grades int
	const (
		A Grades = iota
		B 
		C 
		D 
		E 
		F
	)
	fmt.Println("Grades A ", A)
	fmt.Println("Grades B ", B)
	fmt.Println("Grades C ", C)
	fmt.Println("Grades D ", D)
	fmt.Println("Grades E ", E)
	fmt.Println("Grades F ", F)
}
