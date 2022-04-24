package main

import "fmt"

func foo(y *int){
	*y = *y + 1
	fmt.Println("*y ", *y)
}

func main() {
	x := 2
	foo(&x)
	fmt.Println("x ", x)
}
