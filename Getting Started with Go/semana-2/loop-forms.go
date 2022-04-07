package main

import "fmt"

func main() {
	for i :=0; i<10; i++ {
		fmt.Println("hi i ", i)
	}
	
	j := 0
	for j < 10 {
		fmt.Println("hi j ", j)
		j++
	}
	
	for {
		fmt.Println("hi ")
	}
}
