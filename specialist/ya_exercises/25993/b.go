package main

import (
	"fmt"
	"regexp"
	"unicode/utf8"
)

type Sequence struct {
	message string
}

func New(newMessage string) *Sequence {
	return &Sequence{newMessage}
}

// Calc max length and return this value
func (s *Sequence) FindMax() int {
	re, _ := regexp.Compile(`о{2,}`)
	res := re.FindAllString(s.message, -1)

	max := 0
	for _, val := range res {
		currentRunes := utf8.RuneCountInString(val)

		if currentRunes > max {
			max = currentRunes
		}
	}

	return max
}

func main() {
	var input string
	fmt.Scan(&input)
	message := Sequence{input}
	fmt.Println(message.FindMax())
}
