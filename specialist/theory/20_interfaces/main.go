package main

import "fmt"

// go run github.com/dozer111/specialist/theory/20_interfaces/main.go

// 1 Декларируем интерфейс
type Figure interface {
	Area() int
	Perimeter() int
}

// 2 декларируем претендентов
// Когда все методы интерфейса реализованы, говорят, что
// CIRCLE УДОВЛЕТВОРЯЕТ УСЛОВИЯМ интерфейса Figure
// Circle реализует интерфейс Figure

// 2.1
type Circle struct {
	square int
}

// 2.2
type Rectangle struct {
	width, height int
}

func (c Circle) Area() int {
	return c.square * c.square * 3
}

func (c Circle) Perimeter() int {
	return c.square * 15
}

func (r Rectangle) Area() int {
	return r.width * r.height
}

func (r Rectangle) Perimeter() int {
	return (r.width + r.height) * 2
}

func main() {
	rect1 := Rectangle{10, 15}
	rect2 := Rectangle{7, 8}
	rect3 := Rectangle{12, 11}

	circle1 := Circle{2}
	circle2 := Circle{5}
	circle3 := Circle{7}

	figures := []Figure{rect1, rect2, rect3, circle1, circle2, circle3}
	result := 0
	for _, figure := range figures {
		result += figure.Area()
	}

	fmt.Printf("Total area is %d\n", result)
}
