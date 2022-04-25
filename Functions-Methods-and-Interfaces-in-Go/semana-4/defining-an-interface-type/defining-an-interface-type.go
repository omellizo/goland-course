package main

import "fmt"

type Speaker interface {
	Speak()
}

type Dog struct {
	name string
}

func (d Dog) Speak() {
	fmt.Println(d.name)
}

func main() {
	var speaker1 Speaker
	dog1 := Dog{"Brian"}
	speaker1 = dog1
	speaker1.Speak()
}
