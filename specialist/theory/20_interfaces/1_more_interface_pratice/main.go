package main

import (
	"fmt"
	"unicode/utf8"
)

// go run github.com/dozer111/specialist/theory/20_interfaces/first/main.go

// 1 интерфейс
type BigWord interface {
	IsBig() bool
}

// 2 претендент, которого надо научить делать IsBig()
type MySuperString string

// 3 реализация IsBig в MySuperString
func (mss MySuperString) IsBig() bool {
	return utf8.RuneCountInString(string(mss)) > 10
}

func main() {
	sample := MySuperString("Test123123123132qsd")
	var word BigWord
	word = sample
	fmt.Println(word, word.IsBig())
}
