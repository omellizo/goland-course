package main

import "fmt"

func main() {
	idMap := map[string]int {"Joe": 123, "Ana": 456, "Lui": 789}
	fmt.Println("idMap ", idMap)

	id, find := idMap["Joe"]
	fmt.Println("id ", id, " find ", find)

	id, find = idMap["Aaa"]
	fmt.Println("id ", id, " find ", find)

	fmt.Println("Len ", len(idMap))
}
