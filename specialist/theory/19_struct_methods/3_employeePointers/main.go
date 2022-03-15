package main

import "fmt"

/*
go run github.com/dozer111/specialist/theory/19_struct_methods/3_employeePointers/main.go

В GO в метод структуры можно

1) передать саму структуру(и она копирнётся в stackFrame)
2) передать указатель на структуру

*/

type EmployeePointer struct {
	name   string
	salary int
}

// valueReceiver
func (e EmployeePointer) setName(newName string) {
	e.name = newName
}

// pointerReceiver
func (e *EmployeePointer) setSalary(newSalary int) {
	e.salary = newSalary
}

func main() {
	employee := EmployeePointer{"Jo", 4000}
	fmt.Println(employee)

	employee.setName("Alex")
	// 2 рабочих варианта как можно через ссылку
	(&employee).setSalary(1200) // 1
	fmt.Println(employee)
	employee.setSalary(1300) // 2
	fmt.Println(employee)
}
