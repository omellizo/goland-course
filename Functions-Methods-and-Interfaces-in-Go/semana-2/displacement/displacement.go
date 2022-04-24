package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func getFloat64(label string, scanner *bufio.Scanner) float64 {
	for {
		fmt.Printf("Enter %s: ", label)
		if scanner.Scan() {
			res, err := strconv.ParseFloat(scanner.Text(), 64)
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

func GenDisplaceFn(acceleration, initialVelocity,  initialDisplacement float64) func (float64) float64 {
	fn := func (time float64) float64 {
		return ((acceleration * math.Pow(time, 2)) / 2) + initialVelocity * time + initialDisplacement
	}
	return fn
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	
	acceleration := getFloat64("Write value to acceleration", scanner)
	initialVelocity := getFloat64("Write value to accinitial velocityeleration", scanner)
	initialDisplacement := getFloat64("Write value to initial displacement", scanner)
	fn := GenDisplaceFn(acceleration, initialVelocity, initialDisplacement)

	fmt.Println(fn(3))
	fmt.Println(fn(5))
}
