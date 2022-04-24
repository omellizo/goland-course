package main

import "fmt"

type MyInt int

func (my MyInt) numDouble () int {
	return int (my * 2)
}

func main() {
	v := MyInt(3)
	fmt.Println(v.numDouble())
}
