package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode/utf8"
)

type SmartList struct {
	words []string
}

func New(newWords []string) *SmartList {
	return &SmartList{newWords}
}

// Sort words and printed it
func (sl *SmartList) GetAnswer() {

	listLen := len(sl.words)

	if listLen == 1 {
		fmt.Println(sl.words[0])
		return
	}

	for i := 0; i < listLen; i++ {
		for ii := i + 1; ii < listLen; ii++ {
			current, next := sl.words[i], sl.words[ii]
			currentCount, nextCount := utf8.RuneCountInString(current), utf8.RuneCountInString(next)

			if (currentCount > nextCount) || (currentCount == nextCount && current > next) {
				sl.words[i] = next
				sl.words[ii] = current
			}

		}
	}

	for _, val := range sl.words {
		fmt.Println(val)
	}

}

func main() {
	var stringsCount int
	fmt.Scan(&stringsCount)

	strings := make([]string, 0, stringsCount)
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < stringsCount; i++ {
		scanner.Scan()
		strings = append(strings, scanner.Text())
	}

	New(strings).GetAnswer()
}
