package main

import (
	"fmt"
	"sort"
	"sync"
)

func main() {
	var size = 12
	fmt.Printf("Please, enter 12 numbers\n")

	var arr = make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Printf("Enter %dth element: \n", i)
		fmt.Scanf("%d", &arr[i])
	}

	var wg sync.WaitGroup
	subarr_len := len(arr) / 4
	subarr1, subarr2, subarr3, subarr4 := arr[:subarr_len], arr[subarr_len:2*subarr_len],
		arr[2*subarr_len:3*subarr_len], arr[3*subarr_len:]
		
	fmt.Println("Your array is: ", arr)
	fmt.Println("subarr1: ", subarr1)
	fmt.Println("subarr2: ", subarr2)
	fmt.Println("subarr3: ", subarr3)
	fmt.Println("subarr4: ", subarr4)

	wg.Add(4)
	go subarrSort(&wg, subarr1)
	go subarrSort(&wg, subarr2)
	go subarrSort(&wg, subarr3)
	go subarrSort(&wg, subarr4)
	wg.Wait()

	fmt.Println("\nSorted subarrays: ")
	fmt.Println("subarr1: ", subarr1)
	fmt.Println("subarr2: ", subarr2)
	fmt.Println("subarr3: ", subarr3)
	fmt.Println("subarr4: ", subarr4)

	sorted := merge(merge(subarr1, subarr2), merge(subarr3, subarr4))

	fmt.Println("\nSorted arr:", sorted)
}

func subarrSort(wg *sync.WaitGroup, subarr []int) {
	sort.Ints(subarr)

	if wg != nil {
		wg.Done()
	}
}

func merge(subarr1, subarr2 []int) []int {
	var len1, len2 = len(subarr1), len(subarr2)
	result := make([]int, len1+len2)
	i, j, k := 0, 0, 0

	for i < len1 && j < len2 {
		if subarr2[j] > subarr1[i] {
			result[k] = subarr1[i]
			i++
		} else {
			result[k] = subarr2[j]
			j++
		}

		k++
	}
	for i < len1 {
		result[k] = subarr1[i]
		i++
		k++
	}
	for j < len2 {
		result[k] = subarr2[j]
		j++
		k++
	}
	return result
}