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
	vslice1 := []int{4, 6, 8, 1, 3, 5}
	fmt.Println(getMax(vslice1...))
	vslice2 := []int{4, 6, 8, 1, 3, 5, 4, 9, 6, 8, 1, 3, 5}
	fmt.Println(getMax(vslice2...))
}
