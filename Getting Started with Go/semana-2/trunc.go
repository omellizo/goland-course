package main

import "fmt"

func main() {
	fmt.Printf("Write a floating point number: ")
	num := 0.0;
	fmt.Scan(&num)
	fmt.Println("Num ", num)

	var numInt = int(num)
	fmt.Println("NumInt ", numInt)
}
