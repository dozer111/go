package main

import "fmt"

// ==================================================================================================================
// В GO есть нюансы при реализации инетефейса
// func (s Struct) => работает со всем
// func (s *Struct) => не работает со всем

type Describer interface {
	Describe()
}

type Student struct {
	name string
	age  int
}

type Professor struct {
	name string
	age  int
}

func (s Student) Describe() {
	fmt.Printf("Student %s %d y.o\n", s.name, s.age)
}

func (p *Professor) Describe() {
	fmt.Printf("Professor %s\n", p.name)
}

func main() {
	// 1. в методе мы работаем с копией
	var desc1 Describer
	student := Student{"Jo St", 24}
	desc1 = student // All gut, ошибки нет
	desc1.Describe()

	student2 := &Student{"Jo2 St", 23}
	desc1 = student2 // тут тоже gut, указатель сам разыменуется и копирнётся
	desc1.Describe()

	// 2 метод работает с указателем
	var desc2 Describer
	professor1 := &Professor{"Dick West", 60}
	desc2 = professor1 // при передаче указателя - All gut
	desc2.Describe()

	// ==================================================================================================================
	// OOOPS, ошибочка
	// Cannot use 'professor2' (type Professor) as the type Describer Type does not implement 'Describer' as the 'Describe' method has a pointer receiver
	// professor2 := Professor{"Dick West2", 61}
	// desc2 = professor2
	// desc2.Describe()

	// Тоесть, если у нас метод структуры требует значение, то можно передать туда И значение, И указатель
	// Если же метод требует указатель, то компилятор сам его НЕ СДЕЛАЕТ!

	// Сам по себе интерфейс - не референсный тип
	// Чтобы этой "ошибочки" не было - в Go очень любят порождающие конструкции, которын
	// возвращают указатели, а не значения
}
