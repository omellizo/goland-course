package main

import "fmt"

func main() {
	idMap := map[string]int {"Joe": 123, "Ana": 456, "Lui": 789}
	fmt.Println("idMap ", idMap)

	for key, val := range idMap {
		fmt.Println(key, val)
	}
}
