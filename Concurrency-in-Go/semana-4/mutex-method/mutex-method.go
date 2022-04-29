package main

import (
	"fmt"
	"sync"
)

var i int = 0
var wg sync.WaitGroup
var mut sync.Mutex

func inc(){
	mut.Lock()
	i++
	mut.Unlock()
	wg.Done()
}

func main(){
	wg.Add(2)
	go inc()
	go inc()
	wg.Wait()
	fmt.Println(i)
}