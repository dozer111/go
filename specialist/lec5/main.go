package main

import (
	"fmt"
	"strings"
)

func main() {

	// условия
	// 1 if/else пример 1
	// 2 if/elseif/else
	// ex1()
	ex2()
}

func ex1() {
	var value int
	fmt.Scan(&value)
	if value%2 == 0 {
		fmt.Println("Чётное")
	} else {
		fmt.Println("Нечетное")
	}
}

func ex2() {
	var color string
	fmt.Scan(&color)

	// 1 вариант через сравнение
	// if color == "red" {
	// 	fmt.Println("RED")
	// } else if color == "green" {
	// 	fmt.Println("GRRR")
	// } else if color == "yellow" {
	// 	fmt.Println("YYYY")
	// } else {
	// 	fmt.Println("Unknown")
	// }

	// второй способ сравнить строки
	if strings.Compare(color, "red") == 0 {
		fmt.Println("RED")
	} else if strings.Compare(color, "green") == 0 {
		fmt.Println("GRRR")
	} else if strings.Compare(color, "yellow") == 0 {
		fmt.Println("YYYY")
	} else {
		fmt.Println("Unknown")
	}
}
