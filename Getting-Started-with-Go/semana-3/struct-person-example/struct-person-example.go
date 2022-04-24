package main

import "fmt"

type Person struct{
	name string
	address string
	phone string
}

func main() {
	var person Person
	person.name = "Oscar"
	person.address = "Calle falsa 123"
	person.phone = "456789123"
	fmt.Println(person)

	var person2 Person
	person2.name = "Viviana"
	person2.address = "Calle cerca 123"
	person2.phone = "98986532"
	fmt.Println(person2)
}