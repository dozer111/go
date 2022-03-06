package main

import (
	"errors"
	"fmt"
)

/*
D. BP: Ферма не для майнинга
https://contest.yandex.ru/contest/25869/problems/D/

Для восстановления сельского хозяйства государство выделяет целевые субсидии на покупку скота начинающим фермерам.
Выделяемая сумма определяется для каждого региона отдельно, при этом устанавливается точное количество голов скота,
которое надо приобрести.
Цены следующие:
бык – 20 тыс. рублей,
корова - 10 тыс. рублей,
теленок – 5 тыс. рублей.
Выделяемую сумму необходимо потратить полностью, иначе финансы сгорят.

Выведите все возможные варианты стада, которое может купить начинающий фермер на эту сумму.
Обратите внимание, что для развития хозяйства необходимо, чтобы в стаде был как минимум один бык.
Гарантируется, что на выделенную сумму можно купить хотя бы один вариант стада, удовлетворяющий всем условиям.
*/

const (
	BULL_AMOUUNT = 20
	COW_AMOUUNT  = 10
	CALF_AMOUUNT = 5
)

func main() {

	var sum, count int

	fmt.Scan(&sum, &count)

	bullsCount := 1
	result := make([][3]int, 0, 20)
	for {
		data, err := t(bullsCount, sum, count)

		if err != nil {
			break
		}

		result = append(result, data)
		bullsCount++
	}

	for _, val := range result {
		fmt.Printf("%d %d %d\n", val[0], val[1], val[2])
	}
}

func t(bullsNow, sum, count int) ([3]int, error) {
	var bulls, cows, calves int = bullsNow, 0, 0

	// 1. просчитываем вариант, когда мы просто упрёмся по деньгам в количество
	maxCalves := (sum - (bulls * BULL_AMOUUNT)) / CALF_AMOUUNT
	calves = maxCalves
	if maxCalves <= count-bulls {
		return [3]int{bulls, cows, calves}, errors.New("Слишком много быков")
	}

	allCount := bulls + calves
	for allCount > count {
		cows++
		calves -= 2
		allCount--
	}

	return [3]int{bulls, cows, calves}, nil
}
