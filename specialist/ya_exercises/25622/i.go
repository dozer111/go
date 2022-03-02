package main

import (
	"fmt"
	"strconv"
)

/*
I. BP: Математик. Часть 1.
https://contest.yandex.ru/contest/25622/problems/I/

Всегда интересно , является ли число простым.
Простые числа - натуральные числа, которые имеют только 2 делителя : 1 и само это число
(т.к. они обязаны отличаться - 1 это не простое число).
Ваша задача вывести все делители числа N.
В случае, если число N - простое, нужно дополнительно сообщить об этом.


Формат вывода:
Все делители числа N в порядке возрастания. Если число N - простое, то еще вывести слово ACHTUNG
*/

func main() {

	var number int

	fmt.Scan(&number)
	if number%2 == 1 {
		fmt.Printf("%d %d\nACHTUNG\n", 1, number)
		return
	}

	dividers := []int{}
	for i := 1; i <= number/2; i++ {
		if number%i == 0 {
			dividers = append(dividers, i)
		}
	}

	var result string = ""
	for _, value := range dividers {
		if value != 1 {
			result += " "
		}
		result += strconv.Itoa(value)

	}

	fmt.Println(result)
}
