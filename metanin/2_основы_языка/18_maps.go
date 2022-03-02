package main

import "fmt"

/*
map - стандартный ассоциативный массив
https://metanit.com/go/tutorial/2.14.php

*/

func main() {
	// ошибка: 	// ошибка:
	// var users map[int]string
	// users[1] = "Alex"
	// users[2] = "Nico"

	var users2 = map[int]string{
		1: "D",
		2: "C",
	}

	fmt.Println(users2)

	fmt.Println(users2[1], users2[2])

	users3 := map[string]int{
		"d": 1,
		"t": 2,
		"q": 3,
	}

	fmt.Println(users3["qq"])

	// проверка на наличие ключа
	if k, ok := users3["qq"]; ok {
		fmt.Println(k)
	} else {
		fmt.Println("Key is invalid")
	}

	users3["Q1"] = 111
	users3["Q2"] = 222
	delete(users3, "Q1")

	fmt.Println(users3)

}
