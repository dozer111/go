package main

import (
	"fmt"
	"strconv"
)

// ==================================================================================================================
// Создаём из нескольких интрефейсов один
//
// go run github.com/dozer111/specialist/theory/21_interfaces/3_embedding/

type Interface1 interface {
	Method1() string
}

type Interface2 interface {
	Method2() string
}

// собираем из 2х интерфейсов третий
type InterfaceM interface {
	Interface1
	Interface2
}

type Student struct {
	name string
	age  int
}

func (s Student) Method1() string {
	return s.name + " " + strconv.Itoa(s.age)
}

func (s Student) Method2() string {
	return strconv.Itoa(s.age) + " => " + s.name
}

func main() {
	var if1 Interface1
	var if2 Interface2
	var ifAll InterfaceM

	// структура ПОЛНОСТЬЮ удовлетворяет все 3 интерфейса
	student := Student{"Jimm", 18}
	if1 = student
	if2 = student
	ifAll = student

	fmt.Println(if1.Method1())
	fmt.Println(if2.Method2())
	fmt.Println(ifAll.Method1())
	fmt.Println(ifAll.Method2())

}
