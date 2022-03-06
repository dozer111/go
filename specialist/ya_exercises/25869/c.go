package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
C. Рамка на даче у бабули
https://contest.yandex.ru/contest/25869/problems/C/



*/

func main() {
	text := getText()

	file, err := os.OpenFile("github.com/dozer111/specialist/ya_exercises/25869/output.txt", os.O_APPEND|os.O_WRONLY, os.ModeAppend)

	defer file.Close()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, val := range text {
		_, err := file.WriteString(val)

		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

func getText() []string {
	file, err := os.Open("github.com/dozer111/specialist/ya_exercises/25869/input.txt")

	if err != nil {
		fmt.Println(err.Error())
	}

	inputInfo := make([]byte, 64)

	defer file.Close()
	input := make([]string, 4)
	for {
		n, err := file.Read(inputInfo)
		if err == io.EOF {
			break
		}

		input = strings.Fields(string(inputInfo[:n]))
	}

	// / высота
	// / ширина
	// / символ

	height, _ := strconv.Atoi(input[0])
	width, _ := strconv.Atoi(input[1])
	symbol := input[2]
	result := make([]string, 0, 20)

	for i := 0; i < height; i++ {

		if i == 0 || i == height-1 {
			result = append(result, printHorizBorder(width, symbol))
		} else {
			result = append(result, printVerticalBorder(width, symbol))
		}
	}
	return result
}

func printHorizBorder(width int, symbol string) string {

	result := ""
	for i := 0; i < width; i++ {
		result += symbol
	}

	return result + "\n"
}

func printVerticalBorder(width int, symbol string) string {
	result := ""
	for i := 0; i < width; i++ {
		if i == 0 || i == width-1 {
			result += symbol
		} else {
			result += " "
		}

	}

	return result + "\n"
}
