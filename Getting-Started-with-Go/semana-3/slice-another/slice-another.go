package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	arr := make([]int, 0, 3)
	for {
		fmt.Println("Enter an integer('X' to quit)")
		var input string
		fmt.Scan(&input)
		if input == "X" {
			fmt.Println("Program has finished")
			break
		}
		value, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		arr = append(arr, value)
		sort.Ints(arr)
		fmt.Println(arr)
	}
}
