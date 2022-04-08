package main

import "fmt"

type Person struct{
	name string
	address string
	phone string
}

func main() {
	p1 := new(Person)
	fmt.Println(p1)

	p2:= Person{name: "Oscar", address: "Calle falsa", phone: "456789"}
	fmt.Println(p2)
}