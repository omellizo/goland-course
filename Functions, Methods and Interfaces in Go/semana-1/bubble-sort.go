package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Swap(array []int, index int) {
	auxNum := array[index]
	array[index] = array[index+1]
	array[index+1] = auxNum
}

func BubbleSort(array []int) {
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-1; j++ {
			if array[j] > array[j+1] {
				Swap(array, j)
			}
		}
	}
}

func getInt(label string, scanner *bufio.Scanner) int {
	for {
		fmt.Printf("Enter %s: ", label)
		if scanner.Scan() {
			res, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("Invalid input.", err)
				continue
			}
			return res
		} else {
			fmt.Println("Invalid input.")
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	array := [10]int{9, 8, 7, 6, 5, -1, 1, 2, 3, 4}
	for i := 0; i < len(array); i++ {
		array[i] = getInt("Write a Integer", scanner)
	}
	fmt.Println("Unordered Array: ", array)
	BubbleSort(array[:])
	fmt.Println("Ordered Array:   ", array)
}
