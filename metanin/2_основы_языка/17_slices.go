package main

import "fmt"

/*
Срезы(slices)
https://metanit.com/go/tutorial/2.13.php

Массив с динамической длиной


*/

func main() {

	// ==================================================================================================================
	// 1 инициализация
	var users []string
	var users2 = []string{"test1", "test2", "test3"}
	users3 := []string{"AAA", "Aaaa", "Axa"}

	fmt.Println(users, users2, users3)

	// ==================================================================================================================
	// 2 доступ к елементам среза такой же, как и к элем массива
	// + пример перебора среза

	fmt.Println(users2[0], users3[1])

	fmt.Println("\n")
	for key, value := range users2 {
		fmt.Printf("KEY => %d, value => %s\n", key, value)
	}

	// ==================================================================================================================
	// создание среза с N пустых значений => make
	var ages []int = make([]int, 15)
	fmt.Println(ages)

	// ==================================================================================================================
	// добавление в срез
	var ages2 []int = make([]int, 2)
	ages2 = append(ages2, 12)
	ages2 = append(ages2, 21, 22, 23)

	fmt.Println(ages2)

	// ==================================================================================================================
	// оператор среза
	initialUsers := [8]string{"Bob", "Alice", "Kate", "Sam", "Tom", "Paul", "Mike", "Robert"}
	userss1 := initialUsers[2:6] // с 3-го по 6-й
	userss2 := initialUsers[:4]  // с 1-го по 4-й
	userss3 := initialUsers[3:]  // с 4-го до конца

	fmt.Println(userss1) // ["Kate", "Sam", "Tom", "Paul"]
	fmt.Println(userss2) // ["Bob", "Alice", "Kate", "Sam"]
	fmt.Println(userss3) // ["Sam", "Tom", "Paul", "Mike", "Robert"]

	// ==================================================================================================================
	// удаление елем со среза
	var friends []string = []string{"Max", "iVAN", "vASSia", "Misha", "Dima"}
	var friends2 []string
	var n = 3
	friends2 = append(friends[:n], friends[n+1:]...)
	fmt.Println(friends, friends2)

}
