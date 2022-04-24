package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

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
	namesMap := make(map[string]string)
	scanner := bufio.NewScanner(os.Stdin)
	name := getInput("name", scanner)
	address := getInput("address", scanner)
	namesMap[name] = address
	barr, err := json.Marshal(namesMap)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}
	fmt.Println(string(barr))
}
