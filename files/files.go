package files

import (
	"archive/zip"
	"bufio"
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// StepNine
/*
Ранее в рамках этого курса при решении задач требовалось прочитать что-то со стандартного ввода
и вывести результат соответственно в стандартный вывод.
При этом кто-то использовал пакет fmt, а кто-то - bufio + os.
Все эти пакеты имеют свои особенности, поэтому в этой задаче мы попробуем решить знакомую нам проблему с помощью пакетов,
которые кто-то мог до этого момента и не применять: bufio + os + strconv.

Задача состоит в следующем: на стандартный ввод подаются целые числа в диапазоне 0-100,
каждое число подается на стандартный ввод с новой строки (количество чисел не известно).
Требуется прочитать все эти числа и вывести в стандартный вывод их сумму.

Несколько подсказок: для чтения вы можете использовать типы bufio.Reader и bufio.
Scanner, а для записи - bufio.Writer. При чтении помните об ошибке io.EOF.
Конвертирование данных из строки в целое число и обратно осуществляется
функциями Atoi и Itoa из пакета strconv соответственно.
Чтение производится из стандартного ввода (os.Stdin), а запись - в стандартный вывод (os.Stdout).
*/
func StepNine() {
	// Все указанные в тексте задачи пакеты уже импортированы (strconv, bufio, os, io)
	// Другие использовать нельзя.

	s := bufio.NewScanner(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	var sum int

	for s.Scan() {
		value, _ := strconv.Atoi(s.Text())

		sum += value
	}

	wr.WriteString(strconv.Itoa(sum) + "\n")
	wr.Flush()
}

// StepThirteen
/*
Поиск файла в заданном формате и его обработка
Данная задача поможет вам разобраться в пакете encoding/csv и path/filepath,
хотя для решения может быть использован также пакет archive/zip
(поскольку файл с заданием предоставляется именно в этом формате).

В тестовом архиве, который вы можете скачать из нашего репозитория на github.com,
https://github.com/semyon-dev/stepik-go/tree/master/work_with_files/task_csv_1
содержится набор папок и файлов. Один из этих файлов является файлом с данными в формате CSV,
прочие же файлы структурированных данных не содержат.

Требуется найти и прочитать этот единственный файл со структурированными данными
(это таблица 10х10, разделителем является запятая), а в качестве ответа необходимо указать число,
находящееся на 5 строке и 3 позиции (индексы 4 и 2 соответственно).
*/
func StepThirteen() {
	zipArchive := "./docs/task.zip"
	reader, err := zip.OpenReader(zipArchive)
	var counterFiles, counterDirs, counterFills int

	if err != nil {
		panic(err)
	}

	defer reader.Close()

	for _, file := range reader.File {
		info := file.FileInfo()

		if !info.IsDir() {
			rc, errRc := file.Open()
			if errRc != nil {
				panic(errRc)
			}

			defer rc.Close()

			bufferFromFile := make([]byte, info.Size())
			length, errBuf := rc.Read(bufferFromFile)
			if errBuf != nil && errBuf != io.EOF {
				panic(errBuf)
			}

			bufferFromFile = bytes.Trim(bufferFromFile, "\x00")
			str := string(bufferFromFile)

			if len(bufferFromFile) > 0 && strings.Contains(str, ",") {
				counterFills++
				// task/dir1/dir2/dir3/dir5/dir5/file1.txt // 42

				buffer := bytes.NewBuffer(bufferFromFile)
				csvReader := csv.NewReader(buffer)

				data, errReader := csvReader.ReadAll()
				if errReader != nil {
					panic(errReader)
				}

				var row, col = 4, 2

				if row < len(data) && col < len(data[row]) {
					result := data[row][col]
					fmt.Printf("length: %d, counterFiles: %d, counterDirs: %d, counterFills: %d, result: %v \n\n",
						length, counterFiles, counterDirs, counterFills, result)
				}
			}

			counterFiles++
		} else {
			counterDirs++
		}
	}

	fmt.Printf("counterFiles: %d, counterDirs: %d, counterFills: %d \n\n", counterFiles, counterDirs, counterFills)
}

// StepFourteen
/*
Поэтапный поиск данных
Данная задача в основном ориентирована на изучение типа bufio.Reader, поскольку этот тип позволяет считывать данные постепенно.

В тестовом файле, который вы можете скачать из нашего репозитория на github.com, содержится длинный ряд чисел, разделенных символом ";". Требуется найти, на какой позиции находится число 0 и указать её в качестве ответа. Требуется вывести именно позицию числа, а не индекс (то-есть порядковый номер, нумерация с 1).

Например:  12;234;6;0;78 :
Правильный ответ будет порядковый номер числа: 4
*/
func StepFourteen() {
	fileName := "./docs/task.data"
	bufferFromFile, errFile := ioutil.ReadFile(fileName)
	if errFile != nil {
		panic(errFile)
	}

	buffer := bytes.NewBuffer(bufferFromFile)
	csvReader := csv.NewReader(buffer)
	csvReader.Comma = ';'

	data, errReader := csvReader.Read()
	if errReader != nil {
		panic(errReader)
	}

	for index, item := range data {
		if len(item) > 0 {
			itemValue, errValue := strconv.Atoi(item)
			if errValue != nil {
				panic(errValue)
			}

			if itemValue == 0 {
				fmt.Printf("index: %d, item: %v, itemValue: %d \n\n", index+1, item, itemValue)
			}
		}
	}
}
