package main

/*
go run github.com/dozer111/specialist/theory/19_struct_methods/figure/main.go
*/

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

type Rectangle struct {
	width, height int
}

func (circle Circle) Perimeter() float64 {
	return circle.radius * 2 * math.Pi
}

func (figure Rectangle) Perimeter() int {
	return (figure.width + figure.height) * 2
}

func main() {
	circle := Circle{10.77}
	rectangle := Rectangle{20, 30}

	fmt.Println(circle.Perimeter())
	fmt.Println(rectangle.Perimeter())
}
