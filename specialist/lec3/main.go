package main

import "fmt"

func main() {

	var (
		name string
		sur  string
		age  int
	)

	// или
	// fmt.Scan(&name)
	// fmt.Scan(&sur)
	// fmt.Scan(&age)
	// или
	fmt.Scan(&name, &sur, &age)

	fmt.Printf("Name: %s\nSur: %s\nAge:%d\n", name, sur, age)
}
