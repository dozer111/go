package main

import "fmt"

/*
A. Тоже ягода
https://contest.yandex.ru/contest/25869/problems/?nc=fp6n4nK4

В один из жарких летних дней Петя и его друг Вася решили купить арбуз.
Они выбрали самый большой и самый спелый, на их взгляд.
После недолгой процедуры взвешивания весы показали w килограмм.
Поспешно прибежав домой, изнемогая от жажды, ребята начали делить приобретенную ягоду,
однако перед ними встала нелегкая задача.
Петя и Вася являются большими поклонниками четных чисел, поэтому хотят поделить арбуз так,
чтобы доля каждого весила именно четное число килограмм, при этом не обязательно, чтобы доли были равными по величине.
Ребята очень сильно устали и хотят скорее приступить к трапезе, поэтому Вы должны подсказать им,
удастся ли поделить арбуз, учитывая их пожелание.
Разумеется, каждому должен достаться кусок положительного веса.

Формат ввода
В первой и единственной строке входных данных записано целое число w (1 < w < 100) — вес купленного ребятами арбуза.

Выведите YES, если ребята смогут поделить арбуз на две части, каждая из которых весит четное число килограмм, и NO в противном случае
*/

func main() {
	var weight int

	fmt.Scan(&weight)

	if weight%2 == 0 && weight != 2 {
		fmt.Println("YES")
		return
	}

	fmt.Println("NO")
}
