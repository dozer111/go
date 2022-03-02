/*
https://contest.yandex.ru/contest/25606/problems/G/

G. Математик из 3-го класса
Напишите программу, выполняющую подсчет периметра и площади прямоугольника.

На вход программе подается 2 числа (длины сторон предполагаемого прямоугольника)

Пример

Ввод
	2
	3

Вывод
	Периметр: 10
	Площадь: 6
*/

package main

import "fmt"

func main() {
	var side1, side2 int

	fmt.Scan(&side1, &side2)

	perim := (side1 + side2) * 2
	radius := (side1 * side2)

	fmt.Printf("Периметр: %d\n", perim)
	fmt.Printf("Площадь: %d", radius)
}
