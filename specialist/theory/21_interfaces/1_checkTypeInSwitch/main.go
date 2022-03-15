package main

import "fmt"

// ==================================================================================================================
// Проверяем, что с типом интерфейса можно точно так же сравниваться
// как и с любым другим типажом языка

type Describer interface {
	Describe()
}

type Student struct {
	name string
	age  int
}

func (s Student) Describe() {
	fmt.Println("Student:", s.name, " =>", s.age)
}

func TypeFinder(value interface{}) {

	// втупую value.Describe() потом сделать yt получится
	//  value.Describe undefined (type interface {} is interface with no methods)
	// switch value.(type) {
	switch v := value.(type) {
	case int:
		fmt.Println("This is INT")
	case string:
		fmt.Println("This is string")
	case Describer:
		fmt.Println("DESCRIBER!")
		v.Describe()
	default:
		fmt.Println("UNKNOWN type")
	}

}

func main() {
	data1 := 1
	data2 := "test"
	data3 := 15.5
	data4 := Student{"alex", 22}

	TypeFinder(data1)
	TypeFinder(data2)
	TypeFinder(data3)
	TypeFinder(data4)
}
