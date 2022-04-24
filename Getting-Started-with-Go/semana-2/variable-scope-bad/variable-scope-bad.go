package main

import "fmt"

func f() {
	var x = 4
	fmt.Printf("%d", x)
}
func g() {
	//fmt.Printf("%d", x) // variable x no puede ser accedida desde g()
}
