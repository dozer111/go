package main

import "fmt"

func main() {
	test1 := []int{}
	fmt.Printf("Arr1 => Len %d|Cap %d\n", len(test1), cap(test1))
	test1 = append(test1, 1)
	test1 = append(test1, 2)
	test1 = append(test1, 3)
	test1 = append(test1, 4)
	fmt.Println(test1)
	fmt.Printf("Arr1 => Len %d|Cap %d\n", len(test1), cap(test1))

	test2 := make([]int, 10, 10)
	fmt.Printf("Arr2 => Len %d|Cap %d\n", len(test2), cap(test2))
	test2 = append(test2, 1)
	test2 = append(test2, 2)
	test2 = append(test2, 3)
	test2 = append(test2, 4)
	test2 = append(test2, 5)
	fmt.Println(test2)
	fmt.Printf("Arr2 => Len %d|Cap %d\n", len(test2), cap(test2))

	test3 := make([]int, 0, 10)
	fmt.Printf("Arr2 => Len %d|Cap %d\n", len(test3), cap(test3))
	test3 = append(test3, 1)
	test3 = append(test3, 2)
	test3 = append(test3, 3)
	test3 = append(test3, 4)
	test3 = append(test3, 5)
	fmt.Println(test3)
	fmt.Printf("Arr3 => Len %d|Cap %d\n", len(test3), cap(test3))

}
