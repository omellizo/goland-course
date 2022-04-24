package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Hello! please enter a string: ")
    scanner.Scan()
	input := scanner.Text()
	input =  strings.ToLower(input)
	if strings.HasPrefix(input, "i") {
		if strings.Contains(input, "a") {
			if strings.HasSuffix(input, "n") {
				fmt.Print("Found!")
				return
			}
		}
	}
	fmt.Print("Not Found!")
  }
  