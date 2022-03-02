package main

import "fmt"

/*
операторы defer, panic
https://metanit.com/go/tutorial/2.18.php

defer - выполняет указанный код в конце программы

panic("Message") - выброс исключения
*/

func main() {

	// как мы увидим из примера defer-ы отработают все, как и в случае
	// с классическим деструктором

	defer fmt.Println("Test")
	defer fmt.Println("Test2")
	defer fmt.Println("Test3")

	fmt.Println("message1")
	fmt.Println("message2")
	panic("Domain exception!")
	fmt.Println("message3")
}
