package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	slice := make([]int, 3)
	fmt.Println("slice:", slice)

	scanner := bufio.NewScanner(os.Stdin)
	index := 0
	for {
		fmt.Print("Write an integer to be added to the slice or X to exit: ")
		scanner.Scan()
		input := scanner.Text()
		if input == "X" || input == "x"{
			break
		} else {
			num, err := strconv.Atoi(input)
			if err == nil {
				if index < 3 {
					slice[index] = num
					index++
					fmt.Println("NOTE: The numbers will not be organized until the fourth iteration")
				}else{
					slice = append(slice, num)
					sort.Ints(slice)	
				}
			}
		}
		fmt.Println(slice)
	}
}
