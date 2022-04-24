package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var a = 2
	var b = 3
	avg := 2 % (a + b)
	avg2 := float64(a+b) / 2
	avg3 := float64(a+b) / 2.0
	avg4 := float64(float64(a+b) / 2.0)
	fmt.Println("avg ", avg)
	fmt.Println("avg2 ", avg2)
	fmt.Println("avg3 ", avg3)
	fmt.Println("avg4 ", avg4)

	f1()
	f2()
	f3()
	f4()
}

func f1() {
	i, _ := strconv.Atoi("10")
	y := i * 2
	fmt.Println(y)
}

func f2() {
	s := strings.Replace("ianianian", "ni", "in", 2)
	fmt.Println(s)
}

func f3() {
	x := 7
	switch {
	case x > 3:
		fmt.Println("1")
	case x > 5:
		fmt.Println("2")
	case x == 7:
		fmt.Println("3")
	default:
		fmt.Println("4")
	}
}

func f4() {
	var xtemp int
	x1 := 0
	x2 := 1
	for x := 0; x < 5; x++ {
		xtemp = x2
		x2 = x2 + x1
		x1 = xtemp
	}
	fmt.Println(x2)
}

/*func f5() {
	var x int
	var y *int
	z := 3
	y = &z
	x = &y
  }*/
