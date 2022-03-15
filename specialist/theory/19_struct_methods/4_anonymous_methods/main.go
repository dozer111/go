package main

import "fmt"

/*
 go run github.com/dozer111/specialist/theory/19_struct_methods/4_anonymous_methods/main.go

*/

type University struct {
	name string
	city string
}

// В структуру Professor встраиваются все поля
type Professor struct {
	fullName string
	age      int
	University
}

func (u *University) GetFullName() string {
	return u.name + " " + u.city
}

func main() {
	professor := Professor{
		fullName: "Jo Gomezz",
		age:      75,
		University: University{
			name: "NAU",
			city: "Kyiv",
		},
	}

	fmt.Println(professor)
	// вызов метода структуры University в структуре Professor
	fmt.Println(professor.GetFullName())
}
