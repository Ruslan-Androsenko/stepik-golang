package conversions

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// StepThree
/*
Напишите функцию с именем convert, которая конвертирует входное число типа int64 в uint16.
Считывать и выводить ничего не нужно!
Считайте что функция main и пакет main уже объявлены!
*/
func StepThree() {
	var number int64 = 98234
	result := convert(number)

	fmt.Println(result)
}

func convert(number int64) uint16 {
	return uint16(number)
}

// StepThirteen
/*
Представьте что вы работаете в большой компании где используется модульная архитектура.
Ваш коллега написал модуль с какой-то логикой (вы не знаете) и передает в вашу программу какие-то данные.
Вы же пишете функцию которая считывает две переменных типа string,
а возвращает число типа int64 которое нужно получить сложением этих строк.

Но не было бы так все просто, ведь ваш коллега не пишет на Go, и он зол из-за того, что гоферам платят больше.
Поэтому он решил подшутить над вами и подсунул вам свинью.
Он придумал вставлять мусор в строки перед тем как вызывать вашу функцию.

Поэтому предварительно вам нужно убрать из них мусор и конвертировать в числа.
Под мусором имеются ввиду лишние символы и спец знаки.
Разрешается использовать только определенные пакеты: fmt, strconv, unicode, strings,  bytes.
Они уже импортированы, вам ничего импортировать не нужно!

Считывать и выводить ничего не нужно!
Ваша функция должна называться adding() !
Считайте что функция и пакет main уже объявлены!

Sample Input:
%^80 hhhhh20&&&&nd

Sample Output:
100
*/
func StepThirteen() {
	first := "%^80"
	second := "hhhhh20&&&&nd"
	result := adding(first, second)

	fmt.Println(result)
}

func adding(first, second string) int64 {
	firstNum := clearNumber(first)
	secondNum := clearNumber(second)

	return firstNum + secondNum
}

func clearNumber(strInput string) int64 {
	var runeStr = []rune(strInput)
	var runeNumber []rune

	for _, val := range runeStr {
		if unicode.IsNumber(val) {
			runeNumber = append(runeNumber, val)
		}
	}

	res, _ := strconv.ParseInt(string(runeNumber), 10, 64)

	return res
}

// StepFourteen
/*
Для решения данной задачи вам понадобится пакет strconv, возможно использовать пакеты strings или encoding/csv,
или даже bufio - вы не ограничены в выборе способа решения задачи.
В решениях мы поделимся своими способами решения этой задачи, предлагаем вам сделать то же самое.

В привычных нам редакторах электронных таблиц присутствует удобное представление числа
с разделителем разрядов в виде пробела, кроме того в России целая часть от дробной отделяется запятой.
Набор таких чисел был экспортирован в формат CSV, где в качестве разделителя используется символ ";".

На стандартный ввод вы получаете 2 таких вещественных числа, в качестве результата требуется вывести
частное от деления первого числа на второе с точностью до четырех знаков после "запятой"
(на самом деле после точки, результат не требуется приводить к исходному формату).

P.S. небольшое отступление, связанное с чтением из стандартного ввода.
Кто-то захочет использовать для этого пакет bufio.Reader.
Это вполне допустимый вариант, но если вы резонно обрабатываете ошибку метода ReadString('\n'),
то получаете ошибку EOF, а точнее (io.EOF - End Of File).

На самом деле это не ошибка, а состояние, означающее, что файл (а os.Stdin является файлом) прочитан до конца.
Чтобы ошибка была обработана правильно, вы можете поступить так:

if err != nil && err != io.EOF {
	...
}
Sample Input:
1 149,6088607594936;1 179,0666666666666

Sample Output:
0.9750
*/
func StepFourteen() {
	inputString, err := bufio.NewReader(os.Stdin).ReadString('\n')
	inputString = strings.Replace(inputString, " ", "", -1)
	inputString = strings.Replace(inputString, "\r\n", "", -1)
	inputString = strings.Replace(inputString, ",", ".", -1)

	if err != nil && err != io.EOF {
		panic(err)
	}

	values := strings.Split(inputString, ";")
	var numbers []float64

	for _, val := range values {
		numberVal, err := strconv.ParseFloat(val, 64)

		if err != nil {
			panic(err)
		}

		numbers = append(numbers, numberVal)
	}

	var quotient float64

	if len(numbers) > 1 {
		quotient = numbers[0] / numbers[1]
	}

	fmt.Println(strconv.FormatFloat(quotient, 'f', 4, 64))
}
