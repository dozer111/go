package main

import "fmt"

// ==================================================================================================================
// ZeroValue интерфейса
//
// go run github.com/dozer111/specialist/theory/21_interfaces/4_zero_value/

// 1 создаём интрефейс
type Rider interface {
	Ride()
	Gas()
	Stop()
}

func main() {
	// 2 смотрим его zeroValue
	var r Rider
	if r == nil {
		fmt.Printf("rider is nil.\nZeroValueType => %T\nZerovalue value => %v\n", r, r)
	}
}
