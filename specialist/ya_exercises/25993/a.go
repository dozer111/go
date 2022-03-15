package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/*
A. Go в старое доброе сочетание
https://contest.yandex.ru/contest/25993/problems/?nc=91BIRLxl




*/

func main() {
	n, m := getParams()
	fmt.Println(combination(n, m))
}

func combination(n int, m int) int {
	return factorial(n) / (factorial(m) * factorial(n-m))
}

func factorial(number int) int {
	result := 1
	for i := 1; i <= number; i++ {
		result *= i
	}

	return result
}

func getParams() (int, int) {
	input := bufio.NewScanner(os.Stdin)

	data := make([]string, 0, 3)
	for {

		if len(data) == 2 {
			break
		}

		if input.Scan() {
			result := input.Text()
			if result == "" {
				break
			}
			data = append(data, result)
		}
	}

	n, _ := strconv.Atoi(data[0])
	m, _ := strconv.Atoi(data[1])

	return n, m
}
