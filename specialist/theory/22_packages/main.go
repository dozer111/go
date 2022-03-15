package main

import "fmt"

// 2 модуля (main/calculator) помещены в 1 пакет main
// Говорят, что модули main И calculator находятся в одном пакете (в одной директории)

// Разделяющая область видимости
// Всё, что в принципе находится внутри данного пакета доступно из любого модуля без импортирования

func main() {
	// Данные функции видны, потому что они реализованы
	// внутри модуля Calculator, который ВХОДИТ В СОСТАВ ПАКЕТА main
	data1 := Add(20, 10)
	data2 := Sub(20, 10)
	data3 := Mult(20, 10)
	data4 := Div(20, 10)

	fmt.Println(data1, data2, data3, data4)

	// Но, теперь нам нужно по-другому собирать приложение
	// Варианты
	// 1. go run main.go calculator.go
	// 2. go build .
	// ./22_packages
	// 3. go install .

	// go mod init <dirName>
}
