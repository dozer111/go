/*
https://contest.yandex.ru/contest/25606/problems/F/

Напишите программу, которая считывает 4 строки - сообщения, а затем еще одну - имя автора.

После чего ваша программа выводит 4 введенные строки в обратном порядке вместе с именем автора на каждой строке.


Пример 1
Ввод
	Мир
	Привет
	Земля
	Луна
	Женя

Вывод

Луна - Женя
Земля - Женя
Привет - Женя
Мир - Женя
*/

package main

import "fmt"

func main() {
	var str1, str2, str3, str4, name string

	fmt.Scan(&str1, &str2, &str3, &str4, &name)

	name = " - " + name
	fmt.Println(str4 + name)
	fmt.Println(str3 + name)
	fmt.Println(str2 + name)
	fmt.Println(str1 + name)
}
