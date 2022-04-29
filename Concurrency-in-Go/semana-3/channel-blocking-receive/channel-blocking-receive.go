package main

import (
	"fmt"
)

func prod(v1, v2 int, c chan int) {
	c <- v1 * v2
}

func main() {
	c := make(chan int)
	go prod(1, 2, c)
	a := <-c
	b := <-c
	fmt.Println(a * b)
}
