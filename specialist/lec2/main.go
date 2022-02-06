package main

import "fmt"

func main() {
	// 1 вывод в консоль
	// fmt.Println("Hello world!")
	// fmt.Println("Hello world!2", "test22")
	// fmt.Print("Hello world!2", "test2")
	// fmt.Printf("\nHello, my name is %s\nAnd sur is %d\n", "Alex", 43)

	// 2 работа с переменными
	// полустрогая типизация - не обязательно указывать, что это за тип данных, компилятор попробует угадать сам
	// var age int
	// var age2 int = 12
	// var age3 = 123

	// fmt.Println(age, age2, age3)
	// fmt.Printf("%T\n", age3)

	// 3 Декларирование и инициализация множества переменных
	// 3.1 одного типа
	var first, second int = 15, 255
	var first2, second2 = 15, 255
	fmt.Printf("First:%d\nSecond:%d\n", first, second)
	fmt.Printf("First:%d\nSecond:%d\n", first2, second2)

	// 3.2 блок переменных
	var (
		myName string = "Alex"
		mySur  string = "Khonko"
		myAge  int    = 25
	)

	fmt.Printf("Name: %s,sur:%s,age:%d\n", myName, mySur, myAge)

	var (
		myName2 = "Alex"
		mySur2  = "Khonko"
		myAge2  = 25
	)

	fmt.Printf("Name: %s,sur:%s,age:%d\n", myName2, mySur2, myAge2)

	// 3.3 короткая декларация := => оператор короткого присваивания
	count := 12
	fmt.Print(count)
}
