package anonymous

import "fmt"

// StepSeven
/*
Используем анонимные функции на практике.

Внутри функции main (объявлять ее не нужно) вы должны объявить функцию вида func(uint) uint,
которую внутри функции main в дальнейшем можно будет вызвать по имени fn.

Функция на вход получает целое положительное число (uint),
возвращает число того же типа в котором отсутствуют нечетные цифры и цифра 0.
Если же получившееся число равно 0, то возвращается число 100.
Цифры в возвращаемом числе имеют тот же порядок, что и в исходном числе.

Пакет main объявлять не нужно!
Вводить и выводить что-либо не нужно!
Уже импортированы - "fmt" и "strconv", другие пакеты импортировать нельзя.
*/
func StepSeven() {
	fn := func(number uint) uint {
		var digits []uint
		var discharge uint = 1

		for number != 0 {
			digit := number % 10

			if digit%2 == 0 && digit != 0 {
				digits = append(digits, digit)
			}

			number /= 10
		}

		for index, val := range digits {
			if index > 0 {
				discharge *= 10
				val *= discharge
			}

			number += val
		}

		if number == 0 {
			number = 100
		}

		return number
	}

	fmt.Println(fn(727178))
}
