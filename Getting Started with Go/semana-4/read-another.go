package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Person struct {
	fname string
	lname string
}

func getInput(label string, scanner *bufio.Scanner) string {
	var res string
	for {
		fmt.Printf("Enter %s: ", label)
		if scanner.Scan() {
			res = scanner.Text()
			break
		} else {
			fmt.Println("Invalid input.")
		}
	}
	return res
}
func main() {
	var data []Person
	stdin_scanner := bufio.NewScanner(os.Stdin)
	filename := getInput("filename", stdin_scanner)
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	file_scanner := bufio.NewScanner(f)
	for file_scanner.Scan() {
		inputs := strings.Split(file_scanner.Text(), " ")
		p := Person{fname: inputs[0], lname: inputs[1]}
		data = append(data, p)
	}
	f.Close()
	for i, p := range data {
		fmt.Printf("Person %d) %s\n", i, p)
	}
}
