package datetimes

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

// вам это понадобится
const now int64 = 1589570165

// StepThree
/*
На стандартный ввод подается строковое представление даты и времени в следующем формате:

1986-04-16T05:20:00+06:00
Ваша задача конвертировать эту строку в Time, а затем вывести в формате UnixDate:

Wed Apr 16 05:20:00 +0600 1986

Для более эффективной работы рекомендуется ознакомиться с документацией о содержащихся в модуле time константах.

Sample Input:
1986-04-16T05:20:00+06:00

Sample Output:
Wed Apr 16 05:20:00 +0600 1986
*/
func StepThree() {
	var inputDT string
	fmt.Scan(&inputDT)

	// fmt.Println(inputDT)

	convertDT, err := time.Parse(time.RFC3339, inputDT)

	if err != nil {
		panic(err)
	}

	// fmt.Println(convertDT)

	outputDT := convertDT.Format(time.UnixDate)

	fmt.Println(outputDT)
}

// StepFour
/*
На стандартный ввод подается строковое представление даты и времени определенного события в следующем формате:

2020-05-15 08:00:00
Если время события до обеда (13-00), то ничего менять не нужно,
достаточно вывести дату на стандартный вывод в том же формате.

Если же событие должно произойти после обеда, необходимо перенести его на то же время на следующий день,
а затем вывести на стандартный вывод в том же формате.

Sample Input:
2020-05-15 08:00:00

Sample Output:
2020-05-15 08:00:00
*/
func StepFour() {
	inputDT, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	inputDT = strings.TrimSpace(inputDT)
	convertDT, err := time.Parse("2006-01-02 15:04:05", inputDT)

	if err != nil {
		panic(err)
	}

	inputHour := convertDT.Hour()

	if inputHour >= 13 {
		convertDT = convertDT.Add(time.Hour * 24)
	}
	fmt.Println(convertDT.Format("2006-01-02 15:04:05"))
}

// StepSeven
/*
На стандартный ввод подается строковое представление двух дат,
разделенных запятой (формат данных смотрите в примере).

Необходимо преобразовать полученные данные в тип Time,
а затем вывести продолжительность периода между меньшей и большей датами.

Sample Input:
13.03.2018 14:00:15,12.03.2018 14:00:15

Sample Output:
24h0m0s
*/
func StepSeven() {
	inputStr, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	inputStr = strings.TrimSpace(inputStr)
	strDates := strings.Split(inputStr, ",")

	if len(strDates) >= 2 {
		layout := "02.01.2006 15:04:05"

		firstDate, err := time.Parse(layout, strDates[0])

		if err != nil {
			panic(err)
		}

		secondDate, err := time.Parse(layout, strDates[1])

		if err != nil {
			panic(err)
		}

		if firstDate.After(secondDate) {
			fmt.Println(firstDate.Sub(secondDate))
		} else {
			fmt.Println(secondDate.Sub(firstDate))
		}
	}
}

// StepEight
/*
На стандартный ввод подаются данные о продолжительности периода (формат приведен в примере).
Кроме того, вам дана дата в формате Unix-Time: 1589570165 в виде константы типа int64
(наносекунды для целей преобразования равны 0).

Требуется считать данные о продолжительности периода, преобразовать их в тип Duration,
а затем вывести (в формате UnixDate) дату и время, получившиеся при добавлении периода к стандартной дате.

Небольшая подсказка: базовую дату необходимо явно перенести в зону UTC с помощью одноименного метода.

Sample Input:
12 мин. 13 сек.

Sample Output:
Fri May 15 19:28:18 UTC 2020
*/
func StepEight() {
	inputStr, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	inputStr = strings.TrimSpace(inputStr)

	var minutes, seconds int
	fmt.Sscanf(inputStr, "%d мин. %d сек.", &minutes, &seconds)

	durationTime := time.Duration(minutes)*time.Minute + time.Duration(seconds)*time.Second
	nowTime := time.Unix(now, 0).UTC()
	resTime := nowTime.Add(durationTime)

	fmt.Println(resTime.UTC().Format("Mon Jan 02 15:04:05 UTC 2006"))
}
