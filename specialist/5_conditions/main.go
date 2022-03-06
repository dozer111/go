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
	// ex2()

	// 3 инициализация в блоке условного оператора
	// отличия от PHP
	// ex3()
	n()
}

func n() {
	var a, b, c int

	fmt.Scan(&a, &b, &c)

	d := b*b - (4 * a * c)

	if d > 0 {
		fmt.Println("два корня")
	} else if d == 0 {
		fmt.Println("один корень")
	} else {
		fmt.Println("корней нет")
	}
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

func ex3() {

	// в go, можно инициализировать значение в рантайме, как и в других языках
	// наподобии for($i = 0;i< 10;$i++)
	// или if($a = $this->getData()){ .... }

	if a := 12; a < 100 {
		fmt.Println("< 100")
	}

	// ошибка. В отличии от PHP переменна после управляющего блока удаляется
	// fmt.Println(a)

}
