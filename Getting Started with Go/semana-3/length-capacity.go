package main

import "fmt"

func main() {
	a1 := [3] string {"a", "b", "c"}
	sli1 := a1[0:1]
	fmt.Println(len(sli1), cap(sli1))
}
