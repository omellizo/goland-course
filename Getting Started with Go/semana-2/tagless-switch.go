package main

import "fmt"

func main() {
	var x = 5;
	switch {
	case  x > 1:
		fmt.Println("case 1")
	case x < -1:
		fmt.Println("case 2")
	default:
		fmt.Println("no case")	
	}

	
}
