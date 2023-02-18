package parallelism

import "fmt"

// StepEight
/*
Напишите функцию которая принимает канал и число N типа int. Необходимо вернуть значение N+1 в канал.
Функция должна называться task().

Внимание! Пакет и функция main уже объявлены, выводить и считывать ничего не нужно!
*/
func StepEight() {
	channel := make(chan int)
	defer close(channel)
	go task(channel, 4)

	result := <-channel
	fmt.Println(result)
}

func task(channel chan int, n int) {
	channel <- n + 1
}

// StepNine
/*
Напишите функцию которая принимает канал и строку.
Необходимо отправить эту же строку в заданный канал 5 раз, добавив к ней пробел.

Функция должна называться task2().
Внимание! Пакет и функция main уже объявлены, выводить и считывать ничего не нужно!
*/
func StepNine() {
	channel := make(chan string)
	defer close(channel)
	go task2(channel, "Hello")

	for i := 0; i < 5; i++ {
		result := <-channel
		fmt.Printf("%s", result)
	}

	fmt.Println()
}

func task2(channel chan string, str string) {
	for i := 0; i < 5; i++ {
		channel <- str + " "
	}
}

// StepTen
/*
Напишите элемент конвейера (функцию), что запоминает предыдущее значение и отправляет значения
на следующий этап конвейера только если оно отличается от того, что пришло ранее.

Ваша функция должна принимать два канала - inputStream и outputStream, в первый вы будете получать строки,
во второй вы должны отправлять значения без повторов. В итоге в outputStream должны остаться значения,
которые не повторяются подряд. Не забудьте закрыть канал ;)

Функция должна называться removeDuplicates()
Выводить или вводить ничего не нужно!
*/
func StepTen() {
	inputStream := make(chan string)
	outputStream := make(chan string)
	defer close(inputStream)
	defer close(outputStream)

	inputStrings := []string{
		"Start", "Test", "One", "Two", "Test", "One", "Three", "Two", "Four", "End",
	}

	for _, str := range inputStrings {
		go removeDuplicates(inputStream, outputStream)
		inputStream <- str
		fmt.Printf("%s, ", <-outputStream)
	}

	fmt.Println()
}

func removeDuplicates(inputStream, outputStream chan string) {
	var oldStrs []string

	test := <-inputStream
	oldStrs = append(oldStrs, test)
	outputStream <- test
}
