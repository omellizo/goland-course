package main

import "fmt"

func getMax(vals ...int) int {
	maxV := 1
	for _, v := range vals {
		if v > maxV {
			maxV = v
		}
	}
	return maxV
}

func main() {
	fmt.Println(getMax(4, 6, 8, 1, 3, 5))
	fmt.Println(getMax(4, 6, 8, 1, 3, 5, 4, 9, 6, 8, 1, 3, 5))
}
