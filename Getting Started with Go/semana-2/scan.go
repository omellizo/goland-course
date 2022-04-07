package main

import "fmt"

func main() {
	var appleNum int
	fmt.Printf("Number of apples?")
	num, err := fmt.Scan(&appleNum)
	fmt.Println("&appleNum ", &appleNum)
	fmt.Println("appleNum ", appleNum)
	fmt.Println("num ", num)
	fmt.Println("err ", err)
	
}
