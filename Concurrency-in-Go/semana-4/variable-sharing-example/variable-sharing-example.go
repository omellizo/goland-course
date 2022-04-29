package main

import (
	"fmt"
	"sync"
)

var i int = 0
var wg sync.WaitGroup

func inc(){
	i++
	wg.Done()
}

func main(){
	wg.Add(2)
	go inc()
	go inc()
	wg.Wait()
	fmt.Println(i)
}