package solutions

import (
	"fmt"
	"math"
	"strconv"
)

// StepTwo
/*
На вход подаются a и b - катеты прямоугольного треугольника.
Нужно найти длину гипотенузы
*/
func StepTwo() {
	var a, b float64
	fmt.Scan(&a, &b)

	c := math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))

	fmt.Println(c)
}

// StepThree
/*
Дана строка, содержащая только английские буквы (большие и маленькие).
Добавить символ ‘*’ (звездочка) между буквами (перед первой буквой и после последней символ ‘*’ добавлять не нужно).

Входные данные
Вводится строка ненулевой длины. Известно также, что длина строки не превышает 1000 знаков.

Выходные данные
Вывести строку, которая получится после добавления символов '*'.

Sample Input:
LItBeoFLcSGBOFQxMHoIuDDWcqcVgkcRoAeocXO

Sample Output:
L*I*t*B*e*o*F*L*c*S*G*B*O*F*Q*x*M*H*o*I*u*D*D*W*c*q*c*V*g*k*c*R*o*A*e*o*c*X*O
*/
func StepThree() {
	var str string
	fmt.Scan(&str)

	var runeStr = []rune(str)
	var newRuneStr []rune
	var length = len(runeStr)

	for i := 0; i < length; i++ {
		newRuneStr = append(newRuneStr, runeStr[i])

		if i < length-1 {
			newRuneStr = append(newRuneStr, rune('*'))
		}
	}

	var newStr = string(newRuneStr)
	fmt.Println(newStr)
}

// StepFour
/*
Дана строка, содержащая только десятичные цифры. Найти и вывести наибольшую цифру.

Входные данные
Вводится строка ненулевой длины. Известно также, что длина строки не превышает
1000 знаков и строка содержит только десятичные цифры.

Выходные данные
Выведите максимальную цифру, которая встречается во введенной строке.
*/
func StepFour() {
	var numStr string
	fmt.Scan(&numStr)

	var runeStr = []rune(numStr)
	var max int

	for i := 0; i < len(runeStr); i++ {
		curent := string(runeStr[i])
		curentNum, err := strconv.Atoi(curent)

		if err == nil && curentNum > max {
			max = curentNum
		}

	}

	fmt.Println(max)
}

// StepFive
/*
На вход подается целое число. Необходимо возвести в квадрат каждую цифру числа и вывести получившееся число.

Например, у нас есть число 9119. Первая цифра - 9. 9 в квадрате - 81.
Дальше 1. Единица в квадрате - 1. В итоге получаем 811181
*/
func StepFive() {
	var strNum, res string
	fmt.Scan(&strNum)

	var runeStr = []rune(strNum)

	for i := 0; i < len(runeStr); i++ {
		curentNum, err := strconv.Atoi(string(runeStr[i]))

		if err == nil {
			res += strconv.Itoa(curentNum * curentNum)
		}
	}

	fmt.Println(res)
}

// StepSeven
/*
Внимательно прочитайте ВСЕ условия и подсказки чтобы правильно решить задачу!
Требуется вычислить период колебаний (t) математического маятника
(мы округлили некоторые значения для удобства проверки),
для этого нужно найти циклическую частоту колебания пружинного маятника (w),
в формуле w встречается масса которую также нужно найти, все нужные формулы приведены ниже:

Напишите три функции, каждая из которых будет выполнять конкретную формулу.
Название функций обязательно должны соответствовать букве формулы: T(), W() и M().
Для того чтобы найти t - необходимо сначала найти w, и т.д.
Так что используйте результат функции W() в формуле функции T() - то-есть вызывайте функцию W() в T().
Аналогично и с W(), M().

t = 6 / w, w = sqrt(k / m), m = p * v

ВАЖНО! Считайте, что пакет main уже объявлен, а также функция main() с вызовом ВАШЕЙ будущей функции T() уже есть.
Несмотря на то, что тестирование будет через ввод-вывод, вам НЕ требуется вводить и выводить что-либо.
Для подсчета используйте УЖЕ ВВЕДЕННЫЕ ГЛОБАЛЬНЫЕ переменные k,p,v ТИПА float64 !!!

Пакет math уже импортирован! Напоминаю: корень (sqrt) можно найти с помощью пакета "math", например:

fmt.Println(math.Sqrt(9))
// результат: 3
*/
func StepSeven() {
	var result = T()
	fmt.Println(result)
}

func T() float64 {
	return 6 / W()
}

func W() float64 {
	k := 1296.0
	return math.Sqrt(k / M())
}

func M() float64 {
	p := 6.0
	v := 6.0

	return p * v
}
