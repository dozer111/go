package main

import "fmt"

/**
go run github.com/dozer111/specialist/17_structures.go


1. Структура -именованый набор полей, определяющий новый тип данных

Структуры бывают
* именованые
* анонимные
*/

// 2 Создаём первые структуры и работаем с ними

type Student struct {
	firstName string
	lastName  string
	age       int
}

type Student2 struct {
	firstName, lastName, middleName string
	age, groupId                    int
}

type Human struct {
	name string
	sur  string
	int
}

func main() {
	// 2
	var student1 Student = Student{
		firstName: "Alex",
		lastName:  "Khonko",
		age:       21,
	}
	fmt.Println(student1)
	printStudent(student1)
	// порядок указания свойств такой же как и в структуре
	student2 := Student{"Jo", "Gigg", 22}
	printStudent(student2)
	student3 := Student{}
	printStudent(student3)

	// 3 Анонимные структуры
	anonStudent := struct {
		nick string
	}{"dozer111"}
	fmt.Println(anonStudent)

	// 4 Иницивлизация пустой структуры
	// способ 1
	student11 := Student{}
	// способ 2
	var student12 Student
	fmt.Println("\n4 Иницивлизация пустой структуры")
	printStudent(student11)
	printStudent(student12)

	// 5 указатели на структуры
	// отображаются не так, как на простых типах ({structData} VS &{structData})
	fmt.Println("\n5 указатели на структуры")
	// вариант 1
	fmt.Println(student1, &student1)
	// вариант 2
	student1Pointer := &student1
	(*student1Pointer).age = 222
	fmt.Println(student1, &student1)
	// вариант 3
	student1Pointer2 := new(Student)
	fmt.Println(student1Pointer2)

	// 6 Доступ к полям через указатель
	fmt.Println("\n6 доступ к полю структуры через указатель")
	fmt.Println("Вариант через (*stuctName).paramName =>", (*student1Pointer).age)
	// Тоже работает, Go умеет разыменовывать на лету
	fmt.Println("Вариант через stuctName.paramName =>", student1Pointer.age)

	// 7 структура с анонимными полями
	human := Human{"alex", "khonko", 22}
	fmt.Println("\n7 Структуры с анонимными полями")
	fmt.Println(
		human,
		human.name,
		human.int, // обращение к анонимному полю
	)
	human.name = "alex2"
	human.int = 1234
	fmt.Println(human)
}

func printStudent(std Student) {
	fmt.Println("===========================")
	fmt.Println("Name:", std.firstName)
	fmt.Println("Sur:", std.lastName)
	fmt.Println("Age:", std.age)
}
