package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewScanner(os.Stdin)

	data := make([]string, 0, 10)
	for {
		if input.Scan() {
			result := input.Text()
			if result == "" {
				break
			}
			data = append(data, result)
		}

	}

	fmt.Println("===========================")
	fmt.Println("data is", data)
	fmt.Println("===========================")
}
