package main

import "fmt"

var PORT = "1234"

const PORT_C = "2222"

// go run github.com/dozer111/specialist/14_constants/main.go
func main() {
	// интересный трюк, я могу в таком варианте по-сути
	// переизменить значение констант

	fmt.Println(PORT)
	PORT = "2222"
	fmt.Println(PORT)
	fmt.Println(PORT_C)
	const PORT_C = "2251"
	fmt.Println(PORT_C)
	// ошибка
	// const PORT_C = "5561"

	// 1 константы имеют область видимости
	// 2 типизированные и нетипизированные константы
	const (
		C1     = 1
		C2 int = 2
	)

	// 3 профит нетипизированных констант
	// в основном лучше всего - целые числа
	// так же можно и float, но не все кейсы заедут
	const NUMERIC = 100 // отработает ВСЁ
	// const NUMERIC = 100.12 // отработает часть преобразований
	// const NUMERIC = "100" // ошибка, нельзя кастануть string в int
	var num1 int = NUMERIC
	var num2 int8 = NUMERIC
	var num3 uint = NUMERIC
	var num4 int64 = NUMERIC
	var num5 float32 = NUMERIC
	var num6 complex64 = NUMERIC

	fmt.Printf("Type => %T, value => %v\n", num1, num1)
	fmt.Printf("Type => %T, value => %v\n", num2, num2)
	fmt.Printf("Type => %T, value => %v\n", num3, num3)
	fmt.Printf("Type => %T, value => %v\n", num4, num4)
	fmt.Printf("Type => %T, value => %v\n", num5, num5)
	fmt.Printf("Type => %T, value => %v\n", num6, num6)

}
