package main

import "fmt"

/*
go run github.com/dozer111/specialist/16_dont_use_array_pointers.go

В GO есть вариант изменять массив через функции с указателями

НО, вместо этого более верное решение - работать со слайсами
*/

func main() {
	// 1 плохо
	// оба варианта рабочие
	arr1 := [3]int{1, 2, 3}
	arr2 := [3]int{2, 3, 4}

	fmt.Println("Before *Array mutate:", arr1, arr2)
	mutateByPointer1(&arr1)
	mutateByPointer2(&arr2)
	fmt.Println("After *Array mutate:", arr1, arr2)

	// 2 более корректный подход
	arr3 := [3]int{21, 22, 23}
	fmt.Println("Before slice mutate:", arr3)
	mutateBySlice(arr3[:])
	fmt.Println("After slice mutate:", arr3)
}

func mutateByPointer1(data *[3]int) {
	(*data)[1] = 15
	(*data)[2] = 16

	// 2.
}

func mutateByPointer2(data *[3]int) {
	// Go сам умеет разыменовывать указатель на массив
	data[1] = 15
	data[2] = 16
}

func mutateBySlice(data []int) {
	data[1] = 444
	data[2] = 225
}
