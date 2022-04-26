/**
** The example of race condition “goal-1000000.go” is a game, the goal is that "sharedNum" be 
** equal to 1.000.000 it starts in 0. But there are 3 routines, the first adds 500 to "sharedNum", 
** the second adds 100 and the third subtracts 30. There are race condition because en the 3 methods 
** use the "sharedNum" variable and the goal is for "shareNum" to be equal to one million, the program 
** at the end show how many times use each method and the time to variable "sharedNum" reach to one millon
**/
package main

import (
	"fmt"
	"time"
)

var sharedNum int = 0
var counter500 = 0
var counter100 = 0
var counter30 = 0

func add500() {
	for {
		sharedNum += 500
		counter500 ++;
		fmt.Println("Plus 500: ", sharedNum)
	}
}

func add100() {
	for {
		sharedNum += 100
		counter100 ++;
		fmt.Println("Plus 100: ", sharedNum)
	}
}

func subtract30() {
	for {
		sharedNum -= 30
		counter30 ++;
		fmt.Println("subtract 50: ", sharedNum)
	}
}

func main() {
	timeStart := time.Now()
	go add500()
	go add100()
	go subtract30()
	
	for sharedNum < 1000000 {
		fmt.Println(".")
	}
	timeEnd := time.Now()

	totalTime := timeEnd.Sub(timeStart)

	fmt.Println("Counter 500: ", counter500)
	fmt.Println("Counter 100: ", counter100)
	fmt.Println("Counter 30:  ", counter30)
	fmt.Println(totalTime.Seconds())
}
