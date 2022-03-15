package main

import (
	"log"
	"testing"
)

// запуск тестов:
// go test
// детальный вывод
// go test -v
// testCoverage
// go test -cover
// go test -cover -v

func TestAdd(t *testing.T) {
	if res := Add(10, 20); res != 30 {
		log.Fatalf("Add(10,30) fail. expected %d, got %d\n", 30, res)
	}
}

func TestSub(t *testing.T) {
	if res := Sub(20, 10); res != 10 {
		log.Fatalf("Sub(20,10) fail. expected %d, got %d\n", 10, res)
	}
}

func TestDiv(t *testing.T) {
	if res := Div(25, 5); res != 5 {
		log.Fatalf("Div(25,5) fail. expected %d, got %d\n", 5, res)
	}
}
