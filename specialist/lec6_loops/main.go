package main

import (
	"fmt"
)

/**
Цикл в GO
https://metanit.com/go/tutorial/2.10.php

*/

func main() {
	// 1 простейший пример for
	// ex1_basic_for()

	// 2 break/continue
	// ex2_break_continue()

	// если нам нужно остановить с внутреннего цикла ещё и внешний(один из внешних)
	// используются
	// 3 лейблы
	// ex3_labels()

	// 4 классический цикл while
	// ex4_basic_while()

	// 5 бесконечтный цикл
	// ex5_infinite_loop()

	f3()
}

func f3() {

	var str string
	for {
		fmt.Scanln(&str)

		if str == "" || str == "\n" {
			fmt.Println("test")
		}
	}
}

func ex1_basic_for() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func ex2_break_continue() {
	fmt.Println("Break code START")
	for i := 0; i < 10; i++ {
		if i > 5 {
			break
		}

		fmt.Println(i)
	}
	fmt.Println("Break code END")

	fmt.Println("\n\nContinue code START")
	for i := 0; i < 20; i++ {
		if i%2 == 0 {
			continue
		}

		fmt.Println(i)
	}
	fmt.Println("Continue code END")
}

func ex3_labels() {
	fmt.Println("Простой вложенный цикл START")
	for i := 0; i < 5; i++ {
		for ii := 1; ii <= i; ii++ {
			fmt.Printf("%d => %d\n", i, ii)
		}
	}
	fmt.Println("Простой вложенный цикл END\n\n")

	fmt.Println("Использование label 1")
outer1:
	for i := 0; i < 5; i++ {
		for ii := 1; ii <= i; ii++ {

			if ii > 2 {
				break outer1
			}

			fmt.Printf("%d => %d\n", i, ii)
		}
	}

	fmt.Println("Использование label 1 END\n\n")

	fmt.Println("Использование label прим 2")
outerEx2:
	for i := 0; i < 10; i++ {
	inner1Ex2:
		for ii := 1; ii <= i; ii++ {
			for iii := ii; iii <= i; iii++ {
				if iii > 5 {
					break inner1Ex2
				}
			}

			if ii > 7 {
				break outerEx2
			}

			fmt.Printf("%d => %d\n", i, ii)
		}
	}

}

func ex4_basic_while() {
	loopVar := 0

	for loopVar < 10 {
		fmt.Printf("WHILE loop %d\n", loopVar)
		loopVar++
	}
}

func ex5_infinite_loop() {

	loop := 0
	for {
		if loop > 20 {
			break
		}

		fmt.Printf("WHILE infinite loop %d\n", loop)
		loop++
	}

	fmt.Printf("WHILE infinite loop %d END\n", loop)
}
