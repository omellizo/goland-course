package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("file.txt")

	barr := []byte{1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5}
	nb, err := f.Write(barr)
	
	fmt.Println("NB ", nb)
	fmt.Println("ERR ", err)

	f.Close()

	ff, erro := os.Open("file.txt")
	fmt.Println("FF ", ff)
	fmt.Println("ERR ", erro)

	barrr := make([]byte, 20)
	fmt.Println("BARRR ", barrr)

	nbb, er := ff.Read(barrr)
	fmt.Println("NBB ", nbb)
	fmt.Println("ER ", er)
	fmt.Println("BARRR ",[]byte(barrr[:nbb]))
	
	ff.Close()
}
