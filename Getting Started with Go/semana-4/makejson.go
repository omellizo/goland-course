package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Write your name: ")
    scanner.Scan()
	name := scanner.Text()
	fmt.Print("Write your address: ")
    scanner.Scan()
	address := scanner.Text()

	dataMap := make(map[string]string)
	fmt.Println("dataMap ", dataMap)
	
	dataMap["name"] = name
	dataMap["address"] =address

	fmt.Println("dataMap ", dataMap)

	var jsonData []byte
	jsonData, err := json.MarshalIndent(dataMap, "", "    ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jsonData))

}
