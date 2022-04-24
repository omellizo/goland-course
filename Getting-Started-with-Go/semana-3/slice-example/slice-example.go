package main

import "fmt"

func main() {
	arr := [...] string {"a", "b", "c", "d", "e", "f", "g"}
	slice1 := arr[1:3]
	slice2 := arr[2:5]
	fmt.Println("Arr ", arr)
	fmt.Println("Slice1 arr[1:3]: ", slice1)
	fmt.Println("Slice2 arr[2:5]: ", slice2)
}
