package main

import "fmt"

func main() {
	var x int = 6
	var y int
	var ip *int
	fmt.Println("X ", x)
	fmt.Println("Y ", y)
	fmt.Println("IP ", ip)

	ip = &x
	y = *ip
	fmt.Println("IP ", ip)
	fmt.Println("Y ", y)

	ptr := new(int)
	fmt.Println("PTR ", ptr)
	fmt.Println("*PTR ", *ptr)
	*ptr = 3
	fmt.Println("PTR ", ptr)
	fmt.Println("*PTR ", *ptr)
}
