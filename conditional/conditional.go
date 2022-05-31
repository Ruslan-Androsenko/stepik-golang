package conditional

import "fmt"

// StepFive
/*
На ввод подается целое число. Если число положительное - вывести сообщение "Число положительное",
если число отрицательное - "Число отрицательное". Если подается ноль - вывести сообщение "Ноль".
Выводить сообщение без кавычек.
*/
func StepFive() {
	var number int
	fmt.Scan(&number)

	if number > 0 {
		fmt.Println("Число положительное")
	} else if number < 0 {
		fmt.Println("Число отрицательное")
	} else {
		fmt.Println("Ноль")
	}
}

// StepSix
/*
По данному трехзначному числу определите, все ли его цифры различны.

Формат входных данных
На вход подается одно натуральное трехзначное число.

Формат выходных данных
Выведите "YES", если все цифры числа различны, в противном случае - "NO".
*/
func StepSix() {
	var number int
	fmt.Scan(&number)

	first := number % 10

	number /= 10
	second := number % 10

	number /= 10
	third := number % 1000

	if first != second && first != third && second != third {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

// StepSeven
/*
Дано неотрицательное целое число. Найдите и выведите первую цифру числа.

Формат входных данных
На вход дается натуральное число, не превосходящее 10000.

Формат выходных данных
Выведите одно целое число - первую цифру заданного числа.
*/
func StepSeven() {
	var number uint
	fmt.Scan(&number)

	if number > 0 && number <= 10000 {
		switch {
		case number == 10000:
			number /= 10000

		case number < 10000 && number >= 1000:
			number /= 1000
		case number < 1000 && number >= 100:
			number /= 100
		case number < 100 && number >= 10:
			number /= 10
		}

		digit := number % 10
		fmt.Println(digit)
	}
}

// StepEight
/*
Определите является ли билет счастливым. Счастливым считается билет,
в шестизначном номере которого сумма первых трёх цифр совпадает с суммой трёх последних.

Формат входных данных
На вход подается номер билета - одно шестизначное  число.

Формат выходных данных
Выведите "YES", если билет счастливый, в противном случае - "NO".
*/
func StepEight() {
	var number uint
	fmt.Scan(&number)

	if number >= 100000 && number < 10000000 {
		sixth := number % 10
		number /= 10

		fifth := number % 10
		number /= 10

		fourth := number % 10
		number /= 10

		third := number % 10
		number /= 10

		second := number % 10
		number /= 10

		first := number % 10

		sumOneTriade := first + second + third
		sumTwoTriade := fourth + fifth + sixth

		if sumOneTriade == sumTwoTriade {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	} else {
		fmt.Println("NO")
	}
}

// StepNine
/*
Требуется определить, является ли данный год високосным, напомним:
Год является високосным если он соответствует хотя бы одному из нижеперечисленных условий:
- кратен 400;
- кратен 4, но не кратен 100.

Входные данные
Вводится единственное число - номер года (целое, положительное, не превышает 10000).

Выходные данные
Требуется вывести слово YES, если год является високосным и NO - в противном случае.
*/
func StepNine() {
	var year uint
	fmt.Scan(&year)

	if year > 0 && year <= 10000 {
		if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	} else {
		fmt.Println("NO")
	}
}
