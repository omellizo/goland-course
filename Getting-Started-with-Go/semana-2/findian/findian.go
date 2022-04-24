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
	strRs, _ := str.ReadString('\n')
	strRstring := strRs[:len(strRs)-2]
	strUpper := strings.ToUpper(strRstring)
	
	if strings.HasPrefix(strUpper, "I") && strings.Contains(strUpper, "A") && strings.HasSuffix(strUpper, "N") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
