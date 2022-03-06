package main

import "fmt"

/*
B. Пирамида из собак
https://contest.yandex.ru/contest/25869/problems/B/

Выведите пирамиду из символов "@" (без кавычек) заданной высоты.

Пример
4

   @
  @@@
 @@@@@
@@@@@@@
*/
func main() {
	var height int

	fmt.Scan(&height)
	if height <= 0 {
		fmt.Println("")
		return
	}

	for i := 1; i <= height; i++ {
		fmt.Println(indent(height, i) + symbs(i))
	}
}

func indent(height int, current int) string {

	indent := height - current
	var result string
	for i := 0; i < indent; i++ {
		result += " "
	}

	return result
}

func symbs(current int) string {
	symbs := (current * 2) - 1

	var result string
	for i := 0; i < symbs; i++ {
		result += "@"
	}

	return result
}
