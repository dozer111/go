package main

import "fmt"

// go run github.com/dozer111/specialist/12_functions/main.go

func main() {
	// 1 вариадический параметр под капотом - slice
	variadic(1, 2, 3)
	variadic(2, 2, 2, 2, 2, 2, 2, 2)
	variadic(5, 6)
	variadic(7)
	variadic()

	// 2 передача slice в вариадический параметр
	sumVariadic(2, 3, 4, 5)
	ints := []int{2, 3, 4, 5, 6}
	sumVariadic(ints...)

	// 3 анонимная ф-я
	add := func(a, b int) int { return a + b }

	fmt.Println(
		add(2, 3),
		add(5, 5),
	)
}

// 1 вариадический параметр под капотом - slice
func variadic(a ...int) {
	fmt.Printf("Type is %T, value is %v\n", a, a)
}

// 2 передача slice в вариадический параметр
func sumVariadic(values ...int) {
	var result int
	for _, val := range values {
		result += val
	}

	fmt.Println(result)
}
