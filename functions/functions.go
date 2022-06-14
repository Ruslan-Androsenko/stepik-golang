package functions

import "fmt"

// StepSeven
/*
Напишите функцию, находящую наименьшее из четырех введённых в этой же функции чисел.

Входные данные
Вводится четыре числа.

Выходные данные
Необходимо вернуть из функции наименьшее из 4-х данных чисел
*/
func StepSeven() {
	result := minimumFromFour()
	fmt.Println("result:", result)
}

func minimumFromFour() int {
	var first, second, third, fourth, result int
	fmt.Scan(&first, &second, &third, &fourth)

	if first < second && first < third && first < fourth {
		result = first
	} else if second < third && second < fourth {
		result = second
	} else if third < fourth {
		result = third
	} else {
		result = fourth
	}

	return result
}

// StepEight
/*
Напишите "функцию голосования", возвращающую то значение (0 или 1),
которое среди значений ее аргументов x, y, z встречается чаще.

Входные данные
Вводится 3 числа - x, y и z (x, y и z равны 0 или 1).

Выходные данные
Необходимо вернуть значение функции от x, y и z.
*/
func StepEight() {
	result := vote(0, 0, 1)
	fmt.Println("result:", result)
}

func vote(x int, y int, z int) int {
	var res int

	if x == y || x == z {
		res = x
	} else {
		res = z
	}

	return res
}

// StepNine
/*
Последовательность Фибоначчи определена следующим образом: φ1=1, φ2=1, φn=φn-1+φn-2 при n>1.
Начало ряда Фибоначчи выглядит следующим образом: 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, ...
Напишите функцию, которая по указанному натуральному n возвращает φn.

Входные данные
Вводится одно число n.

Выходные данные
Необходимо вывести  значение φn.
*/
func StepNine() {
	result := fibonacci(4)
	fmt.Println("result:", result)
}

func fibonacci(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	return fibonacci(n-1) + fibonacci(n-2)
}

// StepEleven
/*
Напишите функцию sumInt, принимающую переменное количество аргументов типа int,
и возвращающую количество полученных функцией аргументов и их сумму.
Пакет "fmt" уже импортирован, функция и пакет main объявлены.

Пример вызова вашей функции:
a, b := sumInt(1, 0)
fmt.Println(a, b)

Результат: 2, 1
*/
func StepEleven() {
	count, sum := sumInt(1, 0, 4, 7, 2, 3)
	fmt.Println("count:", count, "sum:", sum)
}

/* ЗАДАНИЕ:
 * Напишите функцию sumInt, получающую переменное число аргументов типа int,
 * и возвращающую количество переданных аргументов и их сумму.
 */
func sumInt(numbers ...int) (int, int) {
	var sum int
	var count = len(numbers)

	for _, value := range numbers {
		sum += value
	}

	return count, sum
}
