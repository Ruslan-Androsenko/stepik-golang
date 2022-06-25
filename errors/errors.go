package errors

import (
	"errors"
	"fmt"
)

// StepFour
/*
Вы должны вызвать функцию divide которая делит два числа, но она возвращает не только результат деления,
но и информацию об ошибке.
В случае какой-либо ошибки вы должны вывести "ошибка", если ошибки нет - результат функции.
Функция divide(a int, b int) (int, error) получает на вход два числа которые нужно поделить
и возвращает результат (int) и ошибку (error). Пакет main уже объявлен, а нужные пакеты уже импортированы!

Не забудьте считать два числа которые необходимо поделить (передать в функцию)!
*/
func StepFour() {
	var a, b int
	fmt.Scan(&a, &b)

	res, err := divide(a, b)

	if err != nil {
		fmt.Println("ошибка")
	} else {
		fmt.Println(res)
	}
}

func divide(a int, b int) (int, error) {
	var result int
	var error error

	if b != 0 {
		result = a / b
	} else {
		error = errors.New("divide by zero")
	}

	return result, error
}
