package main

import "fmt"

/*
go run github.com/dozer111/specialist/19_struct_methods.go

*/

type Employee struct {
	name     string
	position string
	salary   uint
	currency string
}

func (employee Employee) DisplayInfo() {
	fmt.Println("Name:", employee.name)
	fmt.Println("Position:", employee.position)
	fmt.Printf("Salary %d %s\n", employee.salary, employee.currency)
}

func main() {
	emp1 := Employee{"Alex", "mid back", 3000, "USD"}
	emp1.DisplayInfo()
}
