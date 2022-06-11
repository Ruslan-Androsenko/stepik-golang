package loops

import "fmt"

// StepTwo
/*
Напишите программу, которая выводит квадраты натуральных чисел от 1 до 10.
Квадрат каждого числа должен выводится в новой строке.
*/
func StepTwo() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i * i)
	}
}

// StepThree
/*
Требуется написать программу, при выполнении которой с клавиатуры считываются два натуральных числа A и B
(каждое не более 100, A < B). Вывести сумму всех чисел от A до B  включительно.
*/
func StepThree() {
	var a, b, sum int

	for fmt.Scanln(&a, &b); a > 100 || b > 100 || a > b; {
	}

	for i := a; i <= b; i++ {
		sum += i
	}

	fmt.Println(sum)
}

// StepFour
/*
Напишите программу, которая в последовательности чисел находит сумму двузначных чисел, кратных 8.
Программа в первой строке получает на вход число n - количество чисел в последовательности,
во второй строке -- n чисел, входящих в данную последовательность.
*/
func StepFour() {
	var n, number, sum int
	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&number)

		if number >= 10 && number < 100 && number%8 == 0 {
			sum += number
		}
	}

	fmt.Println(sum)
}

// StepFive
/*
Последовательность состоит из натуральных чисел и завершается числом 0.
Определите количество элементов этой последовательности, которые равны ее наибольшему элементу.

Формат входных данных
Вводится непустая последовательность натуральных чисел, оканчивающаяся числом 0
(само число 0 в последовательность не входит, а служит как признак ее окончания).

Формат выходных данных
Выведите ответ на задачу.
*/
func StepFive() {
	var number int
	fmt.Scanln(&number)
	var counter int = 1
	var max = number

	for number != 0 {
		fmt.Scanln(&number)

		if number > max {
			max = number
			counter = 1
		} else if number == max {
			counter++
		}
	}

	fmt.Println(counter)
}

// StepSeven
/*
Найдите первое число от 1 до n включительно, кратное c, но НЕ кратное d.

Входные данные
Вводится 3 натуральных числа n, c, d, каждое из которых не превышает 10000.

Выходные данные
Вывести первое число от 1 до n включительно, кратное c, но НЕ кратное d.
Если такого числа нет - выводить ничего не нужно.
*/
func StepSeven() {
	var n, c, d int
	fmt.Scan(&n, &c, &d)

	for n > 10000 || c > 10000 || d > 10000 || c <= 0 || d <= 0 {
		fmt.Scan(&n, &c, &d)
	}

	for i := 1; i <= n; i++ {
		if i%c == 0 && i%d != 0 {
			fmt.Println(i)
			break
		}
	}
}

// StepEight
/*
Напишите программу, которая считывает целые числа с консоли по одному числу в строке.

Для каждого введённого числа проверить:

если число меньше 10, то пропускаем это число;
если число больше 100, то прекращаем считывать числа;
в остальных случаях вывести это число обратно на консоль в отдельной строке.
*/
func StepEight() {
	var number int

	for {
		fmt.Scanln(&number)

		if number < 10 {
			continue
		} else if number > 100 {
			break
		} else {
			fmt.Println(number)
		}
	}
}

// StepNine
/*
Вклад в банке составляет x рублей. Ежегодно он увеличивается на p процентов,
после чего дробная часть копеек отбрасывается.
Каждый год сумма вклада становится больше.
Определите, через сколько лет вклад составит не менее y рублей.

Входные данные
Программа получает на вход три натуральных числа: x, p, y.

Выходные данные
Программа должна вывести одно целое число.
*/
func StepNine() {
	var deposit, percent, profit int
	year := 1
	var value float64

	fmt.Scan(&deposit, &percent, &profit)

	for {
		value = float64(percent) / 100.0 * float64(deposit)
		deposit += int(value)

		if deposit >= profit {
			break
		} else {
			year++
		}
	}

	fmt.Println(year)
}

// StepTen
/*
Даны два числа. Определить цифры, входящие в запись как первого, так и второго числа.

Входные данные
Программа получает на вход два числа. Гарантируется, что цифры в числах не повторяются.
Числа в пределах от 0 до 10000.

Выходные данные
Программа должна вывести цифры, которые имеются в обоих числах, через пробел.

Цифры выводятся в порядке их нахождения в первом числе.
Внимание! Если вам сложно решить эту задачу, пропустите и изучайте материал дальше,
затем вернетесь к этой задаче позже.
*/
func StepTen() {
	var first, second uint
	var intersect []uint

	for first <= 0 || second <= 0 || first > 10000 || second > 10000 {
		fmt.Scan(&first, &second)
	}

	for first > 0 {
		digitFirst := first % 10
		tempSecond := second

		for tempSecond > 0 {
			digitSecond := tempSecond % 10

			if digitFirst == digitSecond {
				intersect = append(intersect, digitFirst)
				break
			}

			tempSecond /= 10
		}

		first /= 10
	}

	for i := len(intersect) - 1; i >= 0; i-- {
		fmt.Print(intersect[i], " ")
	}
}
