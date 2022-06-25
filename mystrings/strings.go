package mystrings

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

// StepSeven
/*
На вход подается строка. Нужно определить, является ли она правильной или нет.
Правильная строка начинается с заглавной буквы и заканчивается точкой.
Если строка правильная - вывести Right иначе - вывести Wrong

Маленькая подсказка: fmt.Scan() считывает строку до первого пробела,
вы можете считать полностью строку за раз с помощью bufio:

text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
*/
func StepSeven() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	text = strings.Trim(text, "\r\n")

	var runeText = []rune(text)
	var isRight = unicode.IsUpper(runeText[0]) && strings.HasSuffix(text, ".")

	if isRight {
		fmt.Println("Right")
	} else {
		fmt.Println("Wrong")
	}
}

// StepEight
/*
На вход подается строка. Нужно определить, является ли она палиндромом.
Если строка является палиндромом - вывести Палиндром иначе - вывести Нет.
Все входные строчки в нижнем регистре.

Палиндром — буквосочетание, слово или текст, одинаково читающееся в обоих направлениях
(например, "топот", "око", "заказ").
*/
func StepEight() {
	var str string
	fmt.Scan(&str)

	var runeStr = []rune(str)
	var reverse []rune

	for i := len(runeStr) - 1; i >= 0; i-- {
		reverse = append(reverse, runeStr[i])
	}

	var newStr = string(reverse)

	if str == newStr {
		fmt.Println("Палиндром")
	} else {
		fmt.Println("Нет")
	}
}

// StepNine
/*
Даются две строки X и S. Нужно найти и вывести первое вхождение подстроки S в строке X.
Если подстроки S нет в строке X - вывести -1
*/
func StepNine() {
	var strX, strS string
	fmt.Scan(&strX, &strS)
	result := strings.Index(strX, strS)

	fmt.Println(result)
}

// StepTen
/*
На вход дается строка, из нее нужно сделать другую строку,
оставив только нечетные символы (считая с нуля)
*/
func StepTen() {
	var str string
	fmt.Scan(&str)

	var runeStr = []rune(str)
	var runeNewStr []rune

	for i := 0; i < len(runeStr); i++ {
		if i%2 != 0 {
			runeNewStr = append(runeNewStr, runeStr[i])
		}
	}

	var newStr = string(runeNewStr)
	fmt.Println(newStr)
}

// StepEleven
/*
Дается строка. Нужно удалить все символы, которые встречаются более одного раза и вывести получившуюся строку
*/
func StepEleven() {
	var str string
	fmt.Scan(&str)

	var runeStr = []rune(str)
	var runeNewStr []rune

	for i := 0; i < len(runeStr); i++ {
		countSubstr := strings.Count(str, string(runeStr[i]))

		if countSubstr == 1 {
			runeNewStr = append(runeNewStr, runeStr[i])
		}
	}

	var newStr = string(runeNewStr)
	fmt.Println(newStr)
}

// StepTwelve
/*
Ваша задача сделать проверку подходит ли пароль вводимый пользователем под заданные требования.
Длина пароля должна быть не менее 5 символов, он должен содержать только цифры и буквы латинского алфавита.
На вход подается строка-пароль. Если пароль соответствует требованиям - вывести "Ok", иначе вывести "Wrong password"
*/
func StepTwelve() {
	var str string
	fmt.Scan(&str)
	var runeStr = []rune(str)
	var isCorrect = true
	var isLessLen = len(runeStr) >= 5

	for _, value := range runeStr {
		if !unicode.Is(unicode.Latin, value) && !unicode.IsDigit(value) && !unicode.IsSpace(value) {
			isCorrect = false
		}
	}

	if isLessLen && isCorrect {
		fmt.Println("Ok")
	} else {
		fmt.Println("Wrong password")
	}
}
