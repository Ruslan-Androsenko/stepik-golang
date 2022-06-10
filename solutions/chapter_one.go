package solutions

import (
	"fmt"
	"math"
)

// StepTwo
/*
Дано трехзначное число. Найдите сумму его цифр.

Формат входных данных
На вход дается трехзначное число.

Формат выходных данных
Выведите одно целое число - сумму цифр введенного числа.
*/
func StepTwo() {
	var number int
	fmt.Scan(&number)

	first := number % 10
	number /= 10

	second := number % 10
	number /= 10

	third := number % 10
	sum := first + second + third

	fmt.Println(sum)
}

// StepThree
/*
Дано трехзначное число. Переверните его, а затем выведите.

Формат входных данных
На вход дается трехзначное число, не оканчивающееся на ноль.

Формат выходных данных
Выведите перевернутое число.
*/
func StepThree() {
	var number int
	fmt.Scan(&number)

	first := number % 10
	number /= 10

	second := number % 10
	number /= 10

	third := number % 10
	res := first*100 + second*10 + third

	fmt.Println(res)
}

// StepFour
/*
Идёт k-я секунда суток. Определите, сколько целых часов h и целых минут m прошло с начала суток.
Например, если
k = 13257 = 3 * 3600 + 40 * 60 + 57,
то h = 3 и m = 40.

Входные данные
На вход программе подается целое число k (0 < k < 86399).

Выходные данные
Выведите на экран фразу:
It is ... hours ... minutes.

Вместо многоточий программа должна выводить значения h и m, отделяя их от слов ровно одним пробелом.
*/
func StepFour() {
	var secondOfDay int

	for secondOfDay <= 0 || secondOfDay > 86399 {
		fmt.Scan(&secondOfDay)
	}

	hours := secondOfDay / 3600
	minutes := (secondOfDay - hours*3600) / 60

	fmt.Printf("It is %d hours %d minutes.", hours, minutes)
}

// StepFive
/*
Заданы три числа - a, b, c (a < b < c) - длины сторон треугольника.
Нужно проверить, является ли треугольник прямоугольным.
Если является, вывести "Прямоугольный". Иначе вывести "Непрямоугольный"
*/
func StepFive() {
	var a, b, c float64
	fmt.Scan(&a, &b, &c)

	// Теорема Пифагора:
	// квадрат гипотенузы - равен сумме
	// квадратов катетов

	c1 := math.Sqrt(a*a + b*b)
	b1 := math.Sqrt(a*a + c*c)
	a1 := math.Sqrt(b*b + c*c)

	if a1 == a || b1 == b || c1 == c {
		fmt.Println("Прямоугольный")
	} else {
		fmt.Println("Непрямоугольный")
	}
}

// StepSix
/*
Входные данные
Даны три натуральных числа a, b, c. Определите, существует ли треугольник с такими сторонами.

Выходные данные
Если треугольник существует, выведите строку "Существует", иначе выведите строку "Не существует".
Строку выводите без кавычек.
*/
func StepSix() {
	var a, b, c float64
	fmt.Scan(&a, &b, &c)

	ab := a + b
	bc := b + c
	ac := a + c

	if c <= ab && a <= bc && b <= ac {
		fmt.Println("Существует")
	} else {
		fmt.Println("Не существует")
	}
}

// StepSeven
/*
Даны два числа. Найти их среднее арифметическое.

Формат входных данных
На вход дается два целых положительных числа a и b.

Формат выходных данных
Программа должна вывести среднее арифметическое чисел a и b (ответ может быть целым числом или дробным)
*/
func StepSeven() {
	var a, b int
	fmt.Scan(&a, &b)

	var res float64 = float64(a+b) / 2.0
	fmt.Println(res)
}

// StepEight
/*
Количество нулей
По данным числам, определите количество чисел, которые равны нулю.

Входные данные
Вводится натуральное число N, а затем N чисел.

Выходные данные
Выведите количество чисел, которые равны нулю.
*/
func StepEight() {
	var n, number, zeroCounter int

	for n <= 0 {
		fmt.Scan(&n)
	}

	for i := 0; i < n; i++ {
		fmt.Scan(&number)

		if number == 0 {
			zeroCounter++
		}
	}

	fmt.Println(zeroCounter)
}

// StepNine
/*
Количество минимумов
Найдите количество минимальных элементов в последовательности.

Входные данные
Вводится натуральное число N, а затем N целых чисел последовательности.

Выходные данные
Выведите количество минимальных элементов последовательности.
*/
func StepNine() {
	var n, counter, number int
	var min int = math.MaxInt
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&number)

		if number < min {
			min = number
			counter = 1
		} else if number == min {
			counter++
		}
	}

	fmt.Println(counter)
}

// StepTen
/*
Цифровой корень
Цифровой корень натурального числа — это цифра, полученная в результате итеративного процесса суммирования цифр,
на каждой итерации которого для подсчета суммы цифр берут результат, полученный на предыдущей итерации.
Этот процесс повторяется до тех пор, пока не будет получена одна цифра.

Например цифровой корень 65536 это 7 , потому что 6+5+5+3+6=25 и 2+5=7 .
По данному числу определите его цифровой корень.

Входные данные
Вводится одно натуральное число n, не превышающее 10^7.

Выходные данные
Вывести цифровой корень числа n.
*/
func StepTen() {
	var max = int(math.Pow(10, 7))
	var number, digitalCore int

	for number == 0 || number > max {
		fmt.Scan(&number)
	}

	var (
		digits []int
		temp   int = number
		k      int
	)

	for {
		digits = make([]int, 0)

		for temp > 0 {
			digits = append(digits, temp%10)

			temp /= 10
			k++
		}

		if len(digits) > 1 {
			temp = 0

			for _, value := range digits {
				temp += value
			}
		} else {
			digitalCore = digits[0]
			break
		}
	}

	fmt.Println(digitalCore)
}

// StepEleven
/*
Самое большое число, кратное 7
Найдите самое большее число на отрезке от a до b, кратное 7 .

Входные данные
Вводится два целых числа a и b (a≤b).

Выходные данные
Найдите самое большее число на отрезке от a до b (отрезок включает в себя числа a и b),
кратное 7 , или выведите "NO" - если таковых нет.
*/
func StepEleven() {
	var a, b int
	var max int = math.MinInt
	fmt.Scan(&a, &b)

	for a > b {
		fmt.Scan(&a, &b)
	}

	for i := a; i <= b; i++ {
		if i%7 == 0 && i > max {
			max = i
		}
	}

	if max > math.MinInt {
		fmt.Println(max)
	} else {
		fmt.Println("NO")
	}
}

// StepTwelve
/*
По данному числу n закончите фразу "На лугу пасется..." одним из возможных продолжений:
"n коров", "n корова", "n коровы", правильно склоняя слово "корова".

Входные данные
Дано число n (0<n<100).

Выходные данные
Программа должна вывести введенное число n и одно из слов (на латинице):
korov, korova или korovy, например, 1 korova, 2 korovy, 5 korov.
Между числом и словом должен стоять ровно один пробел.
*/
func StepTwelve() {
	var n int
	var cows string

	for n <= 0 || n > 100 {
		fmt.Scan(&n)
	}

	last := n % 10

	switch {
	case last == 0,
		n >= 5 && n <= 19,
		last >= 5 && last <= 9:
		cows = "korov"

	case n == 1 || last == 1:
		cows = "korova"

	case n == 2 || last == 2,
		n == 3 || last == 3,
		n == 4 || last == 4:
		cows = "korovy"
	}

	fmt.Printf("%d %s", n, cows)
}

// StepThirteen
/*
По данному числу N распечатайте все целые значения степени двойки, не превосходящие N, в порядке возрастания.

Входные данные
Вводится натуральное число.

Выходные данные
Выведите ответ на задачу.
*/
func StepThirteen() {
	var n, res int
	fmt.Scan(&n)

	for i := 0; ; i++ {
		res = int(math.Pow(2, float64(i)))

		if res > n {
			break
		}

		fmt.Print(res, " ")
	}
}

// StepFourteen
/*
Номер числа Фибоначчи
Дано натуральное число A > 1. Определите, каким по счету числом Фибоначчи оно является,
то есть выведите такое число n, что Fn=A. Если А не является числом Фибоначчи, выведите число -1.

Входные данные
Вводится натуральное число.

Выходные данные
Выведите ответ на задачу.
*/
func StepFourteen() {
	var a uint
	var n int = -1
	fibs := []uint{0, 1}

	for a <= 1 {
		fmt.Scan(&a)
	}

	for i := 2; ; i++ {
		first := fibs[i-1]
		second := fibs[i-2]

		curent := first + second
		fibs = append(fibs, curent)

		if a == curent {
			n = i
			break
		}

		if a < curent {
			break
		}
	}

	fmt.Println(n)
}

// StepFifteen
/*
Двоичная запись
Дано натуральное число N. Выведите его представление в двоичном виде.

Входные данные
Задано единственное число N

Выходные данные
Необходимо вывести требуемое представление числа N.
*/
func StepFifteen() {
	var n int
	fmt.Scan(&n)
	fmt.Printf("%b", n)
}

// StepSixteen
/*
Из натурального числа удалить заданную цифру.

Входные данные
Вводятся натуральное число и цифра, которую нужно удалить.

Выходные данные
Вывести число без заданных цифр.
*/
func StepSixteen() {
	var number, digit, result int
	var slice []int
	fmt.Scan(&number, &digit)

	for number > 0 {
		remainder := number % 10

		if remainder != digit {
			slice = append(slice, remainder)
		}

		number /= 10
	}

	for index, value := range slice {
		coefficient := int(math.Pow(10.0, float64(index)))
		result += value * coefficient
	}

	fmt.Println(result)
}
