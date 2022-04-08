package main

import "fmt"

func main() {
	var idMap map[string]int
	idMap = make(map[string]int)
	fmt.Println("idMap ", idMap)

	idMap2 := map[string]int {"Joe": 123}
	fmt.Println("idMap2 ", idMap2)

	fmt.Println(idMap2["Joe"])

	idMap2["Jane"] = 456
	fmt.Println("idMap2 ", idMap2)

	idMap2["Jane"] = 789
	fmt.Println("idMap2 ", idMap2)

	delete(idMap2, "Joe")
	fmt.Println("idMap2 ", idMap2)
}
