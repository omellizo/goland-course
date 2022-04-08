package main

import "fmt"

func main() {
	array := [3] int {1, 2, 3}
	for index, value := range array {
		fmt.Printf("index %d, value %d \n", index, value)
	}
}
