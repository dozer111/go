package main

import "fmt"

// go run github.com/dozer111/specialist/15_pointers/main.go
func main() {
	// 1. Указатель - переменная, которая хранит в себе адрес в памяти другой переменной

	// 2. определение указателя
	var variable int = 30
	var pointer *int = &variable // & => операция взятия адреса в памяти

	fmt.Printf("Type of pointer: %T\n", pointer)
	fmt.Printf("Value of pointer: %v\n", pointer) // адрес в памяти на переменную variable
	fmt.Println(variable, &variable)

	// 3. нулевое значение для указателя
	var zeroPointer *int // zeroValue => nil, т.е это указатель ВНИКУДА
	fmt.Printf("Type => %T, value => %v\n", zeroPointer, zeroPointer)

	// 4. Pointer на pointer
	// бред, но, возможно
	pointerToPointer := pointer
	fmt.Println(pointerToPointer, &pointerToPointer)

	// 5 разыминовывание указателя (получение значения оригинала)
	var numericValue int = 32
	var numVPointer *int = &numericValue
	fmt.Println("Разыминовывание указателя")
	fmt.Printf("Адресс: %v, значение: %v\n", numVPointer, *numVPointer)

	// 6 создание указателей на zeroValue типа
	var num6V int
	var num6Ex1 *int = &num6V // вариант 1
	num6Ex2 := &num6V         // вариант 2
	num6Ex3 := new(int)       // вариант 3 через new

	fmt.Println("Pointer with zeroValue of type")
	fmt.Println(*num6Ex1, *num6Ex2, *num6Ex3)

	// 7 изменение оригинального значения через указатель
	pointerToNum := new(int)
	fmt.Println("Изменяем оригинальное значение через указатель")
	fmt.Printf("Адрес: %v, значение %v\n", pointerToNum, *pointerToNum)
	*pointerToNum = 123
	fmt.Printf("Адрес: %v, значение %v\n", pointerToNum, *pointerToNum)

	// 8 передача указателей в функции
	// КОЛОССАЛЬНЫЙ прирост производительности за счёт того, что мы не копируем данные
	// а работаем напрямую с ними
	var pointerFuncValue int
	add1(&pointerFuncValue)
	add1(&pointerFuncValue)
	add1(&pointerFuncValue)
	add1(&pointerFuncValue)
	add1(&pointerFuncValue)
	fmt.Println(pointerFuncValue)

	// небольшие тесты на слайсах
	d := []int{1, 2}
	q(d)
	fmt.Println(d)
	d2 := []int{1, 2}
	q2(&d2)
	fmt.Println(d2)

	// 9 возврат поинтера из функции
	int1, int2 := getInt(), getInt()
	fmt.Println("9 возврат указателя из функции")
	fmt.Printf("Int1 Type %T, address %v, value %v\n", int1, int1, *int1)
	fmt.Printf("Int2 Type %T, address %v, value %v\n", int2, int2, *int2)
	*int1 = 15
	fmt.Println("=====")
	fmt.Printf("Int1 Type %T, address %v, value %v\n", int1, int1, *int1)
	fmt.Printf("Int2 Type %T, address %v, value %v\n", int2, int2, *int2)
}

func getInt() *int {
	var data int = 123
	return &data
}

func q(data []int) {
	data = append(data, 2)
}

func q2(data *[]int) {
	*data = append(*data, 2)
}

func add1(value *int) {
	*value++
}
