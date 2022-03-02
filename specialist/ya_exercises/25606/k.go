package main

import (
	"fmt"
	"strings"
)

/**
https://contest.yandex.ru/contest/25606/problems/K/
K. Текстовый анализатор. Начало

Аркадий пишет текстовый анализатор комментариев в свой сервис "ВПролете" .
Анализатор работает следующим образом : если пользователь вводит произвольную строку
и в данной строке встречается подстрока "чат" требуется вывести "БОТ",
и "НЕ БОТ" в противном случае.

*/

func main() {
	var message string

	fmt.Scan(&message)

	if strings.Contains(message, "чат") {
		fmt.Println("БОТ")
	} else {
		fmt.Println("НЕ БОТ")
	}
}
