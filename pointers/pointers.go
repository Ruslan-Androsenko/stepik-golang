package pointers

import "fmt"

// StepFive
/*
Напишите функцию, которая умножает значения на которые ссылаются два указателя
и выводит получившееся произведение в консоль.
*/
func StepFive() {
	var x1, x2 int

	fmt.Print("Введите два числа: ")
	fmt.Scan(&x1, &x2)

	testFive(&x1, &x2)
}

func testFive(x1 *int, x2 *int) {
	res := *x1 * *x2
	fmt.Println(res)
}

// StepSix
/*
Поменяйте местами значения переменных на которые ссылаются указатели.
После этого переменные нужно вывести.
*/
func StepSix() {
	var x1, x2 int

	fmt.Print("Введите два числа: ")
	fmt.Scan(&x1, &x2)

	testSix(&x1, &x2)
}

func testSix(x1 *int, x2 *int) {
	oldX1 := *x1
	*x1 = *x2
	*x2 = oldX1

	fmt.Println(*x1, *x2)
}
