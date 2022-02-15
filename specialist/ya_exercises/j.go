package main

import "fmt"

/**
https://contest.yandex.ru/contest/25606/problems/J/
J. Чётные/нечётные

Пользователь вводит 2 вещественных числа, таких, что в сумме они дают целое число

Требуется написать программу, которая определяет,
является ли сумма этих чисел чётным или нечётным числом.
В случае, если сумма чётная - вывести "ЧЁТНОЕ" ,
в противном случае - "НЕЧЁТНОЕ".
*/

func main() {
	var num1, num2 float32

	fmt.Scan(&num1, &num2)

	sum := num1 + num2
	if int32(sum)%2 == 0 {
		fmt.Print("ЧЁТНОЕ")
	} else {
		fmt.Print("НЕЧЁТНОЕ")
	}
}
