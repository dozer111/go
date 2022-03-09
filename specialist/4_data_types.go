/*
Типы данных


*/

package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
	"unsafe"
)

func main() {
	// 1 bool
	bools()

	// 2 numerics, integers
	// int, int8, int16, int32, int64
	// uint, uint8, uint16, uint32, uint64
	// + - * / %
	integers()

	// 3 numerics float
	// float32, float64
	// + - * /
	floats()

	// 4 комплексные числа
	complex_()

	// 5 Strings основы
	// 5.1 strings strlen
	// 5.2 strings поиск подстроки
	// 5.3 углубляемся в руны
	str()
	strlen()
	substr()
	rune_()
}

func bools() {
	var firstBool bool
	fmt.Println(firstBool)
	// boolean operands
	aBoolean, bBoolean := true, false
	fmt.Println("AND: ", aBoolean && bBoolean)
	fmt.Println("OR: ", aBoolean || bBoolean)
	fmt.Println("NOT: ", !aBoolean)
}

func integers() {
	var a int = 32
	b := 92
	fmt.Println("A:", a, "B:", b, "Sum:", a+b)
	// Вывод типа (узнаём, какая же именно разрядность в инта)
	fmt.Printf("Type is %T\n", a)
	fmt.Printf("Type %T has size %d bytes\n", a, unsafe.Sizeof(a)) // 8 bytes === 64 bits => int64

	// var u uint = -12 //Error: constant -12 overflows uint
	// fmt.Println(u)
}

func floats() {
	floatFirst, floatSecond := 5.55, 12.1234
	fmt.Printf("Type1 is %T , type2 is %T \n", floatFirst, floatSecond)

	sum := floatFirst + floatSecond
	sub := floatFirst - floatSecond

	fmt.Println("Sum:", sum, "Sub:", sub)         // не очень хорошый вывод, может быть много чисел после точки
	fmt.Printf("Sum: %.2f Sub: %.2f\n", sum, sub) // так лучше
}

func complex_() {
	c1 := complex(5, 7) // способ задания 1
	c2 := 12 + 15i      // способ задания 2
	fmt.Println(c1, c2)
	fmt.Println(c1 + c2)
	fmt.Println(c1 * c2)
}

func str() {
	name := "alex"
	sur := "khonko"
	fio := name + " " + sur
	fmt.Println(fio)
}

func strlen() {
	name := "alex"
	nameRu := "саня"
	// strlen
	fmt.Println("Length of string", name, len(name))
	fmt.Println("Length of string", name, len(nameRu))
	// mb_strlen
	fmt.Println("MB_ Length of string", nameRu, utf8.RuneCountInString(nameRu))
	// в go нет чаров, символы хранятся в рунах
	// руна - 1 целый utf-ный символ
	// руна - алиас над int32
}

func substr() {
	haystack, needle := "test message", "ess"
	haystack2, needle2 := "вася пупкин", "ас"

	fmt.Println(strings.Contains(haystack, needle))
	fmt.Println(strings.Contains(haystack, "ESS"))
	fmt.Println(strings.Contains(haystack2, needle2))
	fmt.Println(strings.Contains(haystack2, "И"))
}

func rune_() {
	// руна - алиас над int32
	// её можно инициализировать и через символ, и через число
	var rune1 rune
	var rune2 rune = 'Q'
	var rune3 rune = 3331

	fmt.Println(rune1) // default значение типа
	fmt.Printf("Rune as char %c\n", rune2)
	fmt.Printf("Rune as char %c\n", rune3)
}
