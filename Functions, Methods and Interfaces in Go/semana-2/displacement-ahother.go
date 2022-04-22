package main

import (
	"fmt"
)

type F float64

func GenDisplaceFn(a,v,s F) func(F) F{
	return func(t F) F{
		return 0.5 * a * t * t + v * t + s
	}
}

func main(){
	var a,v,s,t F
	fmt.Println("Enter acceleration")
	fmt.Scanf("%f\n",&a)
	fmt.Println("Enter initial velocity")
	fmt.Scanf("%f\n",&v)
	fmt.Println("Enter initial displacement")
	fmt.Scanf("%f\n",&s)
	fmt.Println("Enter time")
	fmt.Scanf("%f\n",&t)
	fn := GenDisplaceFn(a, v, s)
	fmt.Println(fn(t))
}