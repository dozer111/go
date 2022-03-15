package main

import "fmt"

// исполняем разный код, в зависимости от переданного типа

func Do(pretendent interface{}) {
	switch pretendent.(type) {
	case string:
		DoString(pretendent.(string))
	case int:
		DoInt(pretendent.(int))
	default:
		fmt.Printf("Нет обработчика для типа %T\n", pretendent)
	}
}

func DoM(pretendent ...interface{}) {
	for _, val := range pretendent {
		Do(val)
	}
}

func DoString(data string) {
	fmt.Println("String execute => ", data)
}

func DoInt(data int) {
	fmt.Println("Integer => ", data)
}

type Player struct {
	nick string
}

func main() {
	data1 := "data1"
	data2 := Player{"dozer111"}
	data3 := 1234

	Do(data1)
	Do(data2)
	Do(data3)

	fmt.Println("===========================")
	DoM(data1, data2, data3)
}
