package main

import (
	"fmt"
)

/*
H. BP: Геймдизигнер
https://contest.yandex.ru/contest/25622/problems/H/

Вы создали игру в жанре шутер. Теперь ваш дизайнер придумал новое неизвестное никому оружее - дробовик.
Известно, что дробовики стреляют дробью (внезапно, правда?).
Ваша задача - рассчитать суммарный урон, наненсенный выстрелом из дробовика.

Сначала вводится количество дробинок.
Затем урон от каждой дробинки.
Урон от каждой дробинки выражается простой дробью, её числитель и знаменатель вводятся на отдельных строках.
*/

func main() {

	var bulletCount int

	fmt.Scan(&bulletCount)
	// numerator and denominator
	var numerator = make([]int, bulletCount, bulletCount)
	var denominator = make([]int, bulletCount, bulletCount)

	numerator[0] = 1
	denominator[0] = 2

	for i := 0; i < bulletCount; i++ {
		var num, dec int
		fmt.Scan(&num, &dec)
		numerator[i] = num
		denominator[i] = dec
	}

	var maxDenominator int = getDenominator(denominator)
	for index, value := range denominator {
		var multiply int = maxDenominator / value
		numerator[index] *= multiply
	}

	var damage int

	for i := 0; i < len(numerator); i++ {
		damage += numerator[i]
	}

	fmt.Printf("%d/%d\n", damage, maxDenominator)
}

func getDenominator(denominators []int) int {
	maxDenominator := max(denominators)
	currentDenominator := maxDenominator

	for i := 0; i < len(denominators); i++ {
	maxDen:
		for true {
			if currentDenominator%denominators[i] != 0 {
				currentDenominator += maxDenominator
			}
			break maxDen
		}
	}
	return currentDenominator
}

func max(data []int) int {
	result := data[0]

	for _, value := range data {
		if value > result {
			result = value
		}
	}

	return result
}
