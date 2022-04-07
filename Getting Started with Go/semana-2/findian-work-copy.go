package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Printf("Write a string: ")

	str := bufio.NewReader(os.Stdin)
	input, _ := str.ReadString('\n')
	fmt.Println(input)
	input2 := strings.Replace(input, " ", "", -1)
	input2 = input2[:len(input2)-2]
	fmt.Println(input2)
	strUpper := strings.ToUpper(input2)
	result := strUpper
	fmt.Println(result)
	fmt.Println(strings.HasPrefix(strUpper, "I"))
	fmt.Println(strings.Contains(strUpper, "A"))
	fmt.Println(strings.HasSuffix(strUpper, "N"))
	fmt.Println(strings.Index(strUpper, "N"))
	fmt.Println(len(strUpper))

	if strings.HasPrefix(strUpper, "I") && strings.Contains(strUpper, "A") && strings.HasSuffix(strUpper, "N") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}

}
