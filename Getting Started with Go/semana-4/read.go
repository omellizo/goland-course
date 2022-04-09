package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Person struct {
	fname string
	lname string
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Write file name: ")
	scanner.Scan()
	fileName := scanner.Text()

	slice := make([]Person, 0, 1)

	file, err := os.Open(fileName)
	if err != nil {
		os.Exit(0)
	}	
	
	for {
		barrr := make([]byte, 43)
		_, er := file.Read(barrr)
		if er == nil {
			fullName := string(barrr)
			fullName = strings.Replace(fullName, "\r", "", -1)
			fname := fullName[:20]
			lname := fullName[21:41]
			person := Person{fname: fname, lname: lname}
			slice = append(slice, person)
		}else{
			break
		}
	}

	for _, person := range slice {
		fmt.Println(person.fname, person.lname)
	}
	
	file.Close()
}
