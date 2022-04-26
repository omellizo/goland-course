/**
** The example of race condition “goroutine-share-num.go” have a shared num between two goroutines, the first use it and the second change it value.
**/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

var sharedNum int = 0

func useValue() {
	for {
		fmt.Println("USE sharedNum: ", sharedNum)
	}
}

func changeValue() {
	for {
		sharedNum = rand.Intn(100000)
		fmt.Println("CHANGE sharedNum: ", sharedNum)
	}
}

func main() {
	go useValue()
	go changeValue()
	time.Sleep(3 * time.Second)
}