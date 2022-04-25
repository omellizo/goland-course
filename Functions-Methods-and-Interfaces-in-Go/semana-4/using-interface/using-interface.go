package main

import "fmt"

type Shape2D interface {
	Area() float64
	Perimeter() float64
}

type Triangle struct {
	name string
}

func (t Triangle) Area() float64 {
	return 0
}

func (t Triangle) Perimeter() float64 {
	return 0
}

type Rectangle struct {
	name string
}

func (t Rectangle) Area() float64 {
	return 0
}

func (t Rectangle) Perimeter() float64 {
	return 0
}

func FitInYard(s Shape2D) bool {
	if s.Area() > 100 && s.Perimeter() > 100 {
		return true
	}
	return false
}

func DrawShape(s Shape2D) {
	rect, ok := s.(Rectangle)
	if ok {
		fmt.Println("Is a Rectangle", rect)
	}
	rectt, okk := s.(Triangle)
	if okk {
		fmt.Println("Is a Triangle", rectt)
	}
}

func DrawShapeSwitch(s Shape2D) {
	switch sh := s.(type) {
	case Rectangle:
		fmt.Println("Is a Rectangle", sh)
	case Triangle:
		fmt.Println("Is a Triangle", sh)
	}
}

func main() {
	triangle := Triangle{"tri1"}
	DrawShape(triangle)
	DrawShapeSwitch(triangle)

	rectangle := Rectangle{"rect1"}
	DrawShape(rectangle)
	DrawShapeSwitch(rectangle)

}
