package main

import (
	"fmt"
	"sync"
)

var on sync.Once
var wg sync.WaitGroup

func setup() {
	fmt.Println("Init")
}

func dostuff() {
	on.Do(setup)
	fmt.Println("Hello")
	wg.Done()
}

func main() {
	wg.Add(2)
	go dostuff()
	go dostuff()
	wg.Wait()
}