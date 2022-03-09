package main

import "fmt"

/*
настраивание и вложенные структуры

go run github.com/dozer111/specialist/18_structs_pro.go

1. Вложенная структура - структура, которая используется как поле внутри другой структуры
2. Встроенные структуры - когда мы добавляем поля одной структуры в поля другой
3. Структуры можно сравнивать ==
4. Структуры сравниваются по полям
5. Если хотя бы одно поле в структуре не сравнимо - вся структура НЕ СРАВНИМА!!!

*/

type University struct {
	name string
}

type Student18 struct {
	name, sur  string
	university University
}

func main() {
	student := Student18{
		"Alex",
		"Khonko",
		University{"NAU"},
	}

	fmt.Println(student)
	fmt.Println(student.name, student.university.name)
	student.university.name = "NAAAAAA U"
	fmt.Println(student.university.name)
}
