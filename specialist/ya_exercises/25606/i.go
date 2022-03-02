package main

import "fmt"

/**
https://contest.yandex.ru/contest/25606/problems/I/
I. Конкатенатор


Юра складывает числа по определенному правилу, допустим, Юре нужно сложить три числа
ABC

тогда ответом будет число, составленное по правилу
CBA


6
2
3
==
326

4
4
1
==
144

*/

func main() {
	var str1, str2, str3 int
	fmt.Scan(&str1, &str2, &str3)
	// вариант сразу
	// fmt.Print(fmt.Sprintf("%v%v%v", str3, str2, str1))

	// но sprintf можно ведь положить и в переменную
	concat := fmt.Sprintf("%v%v%v", str3, str2, str1)
	fmt.Print(concat)
}
