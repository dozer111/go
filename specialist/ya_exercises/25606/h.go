package main

import "fmt"

/**
H. Хитрая сумма

Напишите программу, которая считывает 2 числа, и считает квадрат суммы этих чисел

(a+b)^2

пример
2
2
==
16

2
3
==
36
*/

func main() {
	var number1, number2 int

	fmt.Scan(&number1, &number2)
	sum := number1 + number2
	fmt.Println(sum * sum)
}
