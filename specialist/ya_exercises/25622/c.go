package main

import "fmt"

/*
https://contest.yandex.ru/contest/25622/problems/C/
C. Полный нуль

Пользователь вводит числа числа одно за другим (каждое с новой строки, все числа ∈ N) до тех пор,
пока не будет введен нуль.

Ваша программа должна выводить вводимые числа до тех пор, пока не будет введен нуль
(Сам нуль выводить не требуется). Никаких арифметических операций над числами производить не требуется.
*/

func main() {

	returnValues := true
	var value int
	for {
		fmt.Scan(&value)

		if value == 0 {
			returnValues = false
		}

		if !returnValues {
			continue
		}

		fmt.Println(value)
	}
}
