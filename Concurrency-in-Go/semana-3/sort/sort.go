package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"sync"
)

func main(){
	slice := make([]int, 12)
	scanner := bufio.NewScanner(os.Stdin)
	
	for i := 0; i < 12; i++ {
		
		slice[i] = getInt(fmt.Sprint("Write a Integer (", (i + 1), "/12)"), scanner)
	}
	fmt.Println(slice)

	sliceSize := len(slice) / 4

	slice1 := slice[:sliceSize]
	slice2 := slice[sliceSize : sliceSize * 2]
	slice3 := slice[sliceSize * 2 : sliceSize * 3]
	slice4 := slice[sliceSize * 3:]

	var wg sync.WaitGroup
	wg.Add(4)

	go sortSlice(&slice1, &wg)
	go sortSlice(&slice2, &wg)
	go sortSlice(&slice3, &wg)
	go sortSlice(&slice4, &wg)

	wg.Wait()

	s1 := append(slice1, slice2...)
	s2 := append(s1, slice3...)
	s3 := append(s2, slice4...)
	sort.Ints(s3)
	fmt.Println("SORTED ", s3)
}

func sortSlice(slice *[]int, wg *sync.WaitGroup){
	fmt.Println("UnOrder ", slice)
	sort.Ints(*slice)
	fmt.Println("InOrder ", slice)
	wg.Done()
}

func getInt(label string, scanner *bufio.Scanner) int {
	for {
		fmt.Printf("Enter %s: ", label)
		if scanner.Scan() {
			res, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("Invalid input.", err)
				continue
			}
			return res
		} else {
			fmt.Println("Invalid input.")
		}
	}
}