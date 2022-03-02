package main

import (
	"fmt"
	"strings"
)

/**
https://contest.yandex.ru/contest/25606/problems/M/
M. Наш отряд

Напишите программу, которая проверяет правильность расчета на "раз два три".

Пользователь вводит последовательно 3 строки.

1. если эти строки "раз", "два", "три" - вывести "ОК" (кириллица)
2. если вместо строки "раз" введена "один" - вывести "ОК" (кириллица)
3. если вместо всех слов введены соответствующие числа "1", "2", "3" - вывести "ОК" (кириллица)
4. "НЕ ПРАВИЛЬНО" - во всех остальных случаях


*/

func main() {
	var str1, str2, str3 string

	fmt.Scan(&str1, &str2, &str3)

	if strings.Contains(str1, "раз") || strings.Contains(str1, "один") {
		if strings.Contains(str2, "два") && strings.Contains(str3, "три") {
			fmt.Println("ОК")
			return
		}
	}

	if str1 == "1" && str2 == "2" && str3 == "3" {
		fmt.Println("ОК")
		return
	}

	fmt.Println("НЕ ПРАВИЛЬНО")
}
