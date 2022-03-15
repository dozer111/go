package main

import "fmt"

// работа с пустыми интерфейсами
// go run github.com/dozer111/specialist/theory/20_interfaces/3_empty_interface/main.go

// бред, можно спокойно заменить на interface{}
type Empty interface {
}

func Describer(pretendent interface{}) {
	fmt.Printf("Type is %T and value is %v\n", pretendent, pretendent)
}

type Player struct {
	nick string
}

func AssertIsInt(pretendent interface{}) {
	// пытаемся привести претендента к типу int
	val, ok := pretendent.(int)
	fmt.Println("Val:", val, "OK?", ok)
}

func main() {
	data1 := "data1"
	data2 := Player{"dozer111"}
	data3 := 1234

	// 1 попрактикуемся
	Describer(data1)
	Describer(data2)
	Describer(data3)

	// 2 проверим на типы
	AssertIsInt(data1)
	AssertIsInt(data2)
	AssertIsInt(data3)
}
