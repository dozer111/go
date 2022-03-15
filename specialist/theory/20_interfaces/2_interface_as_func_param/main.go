package main

import "fmt"

// Интерфейс как параметр функции

type Worker interface {
	Work()
}

type Employee struct {
	name string
	age  int
}

func (e Employee) Work() {
	fmt.Println("Employee", e.name, "is working")
}

func Describe(w Worker) {
	fmt.Printf("Type is %T and value %v\n", w, w)
}

func main() {
	emp := Employee{"Jo", 44}
	Describe(emp)

	var worker Worker
	worker = emp
	Describe(worker)
}
