package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	dat := []byte("Hello, World")
	err := ioutil.WriteFile("outfile.txt", dat, 0777)
	fmt.Println("ERR ", err)

	read, e := ioutil.ReadFile("outfile.txt") 
	fmt.Println("E ", e)
    fmt.Print(string(read))
}
