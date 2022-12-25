package jsons

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// StepSix
/*
На стандартный ввод подаются данные о студентах университетской группы в формате JSON:
{
    "ID":134,
    "Number":"ИЛМ-1274",
    "Year":2,
    "Students":[
        {
            "LastName":"Вещий",
            "FirstName":"Лифон",
            "MiddleName":"Вениаминович",
            "Birthday":"4апреля1970года",
            "Address":"632432,г.Тобольск,ул.Киевская,дом6,квартира23",
            "Phone":"+7(948)709-47-24",
            "Rating":[1,2,3]
        },
        {
            // ...
        }
    ]
}
В сведениях о каждом студенте содержится информация о полученных им оценках (Rating). Требуется прочитать данные,
и рассчитать среднее количество оценок, полученное студентами группы.

Ответ на задачу требуется записать на стандартный вывод в формате JSON в следующей форме:
{
    "Average": 14.1
}
Как вы понимаете, для декодирования используется функция Unmarshal, а для кодирования MarshalIndent
(префикс - пустая строка, отступ - 4 пробела).

Если у вас возникли проблемы с чтением / записью данных, то этот комментарий для вас:
в уроках об интерфейсах и работе с файлами мы рассказывали, что стандартный ввод / вывод - это файлы,
к которым можно обратиться через os.Stdin и os.Stdout соответственно,
они удовлетворяют интерфейсам io.Reader и io.Writer, из них можно читать, в них можно писать.

Один из способов чтения, использовать ioutil.ReadAll:

data, err := ioutil.ReadAll(os.Stdin)
if err != nil {
    // ...
}

// data - тип []byte
Задачу можно выполнить и другими способами, в частности использовать bufio.
Буквально в следующем шаге (через один, на самом деле) будет рассказано о еще одном способе чтения / записи,
можете забежать немного вперед, а затем вернуться к задаче.
*/
func StepSix() {
	//fileName := "./docs/students.json"
	//data, err := ioutil.ReadFile(fileName)
	data, err := ioutil.ReadAll(os.Stdin)

	if err != nil {
		panic(err)
	}

	var groupData Group
	err = json.Unmarshal(data, &groupData)

	if err != nil {
		panic(err)
	}

	var counterRating, counterStudents, sumRating int

	for _, student := range groupData.Students {
		counterStudents++

		for _, rating := range student.Rating {
			sumRating += rating
			counterRating++
		}
	}

	averageRating := float64(counterRating) / float64(counterStudents)
	resultObj := Response{
		Average: averageRating,
	}

	buffer, errIndent := json.MarshalIndent(resultObj, "", "    ")

	if errIndent != nil {
		fmt.Println(errIndent)
		return
	}

	fmt.Println(string(buffer))
}

type Group struct {
	ID       int
	Number   string
	Year     int
	Students []Student
}

type Student struct {
	LastName   string
	FirstName  string
	MiddleName string
	Birthday   string
	Address    string
	Phone      string
	Rating     []int
}

type Response struct {
	Average float64
}

// StepNine
/*
Данная задача ориентирована на реальную работу с данными в формате JSON.
Для решения мы будем использовать справочник ОКВЭД (Общероссийский классификатор видов экономической деятельности),
который был размещен на web-портале data.gov.ru.

Необходимая вам информация о структуре данных содержится в файле structure-20190514T0000.json,
а сами данные, которые вам потребуется декодировать, содержатся в файле data-20190514T0100.json.
Файлы размещены в нашем репозитории на github.com.

Для того, чтобы показать, что вы действительно смогли декодировать документ вам необходимо
в качестве ответа записать сумму полей global_id всех элементов, закодированных в файле.
*/
func StepNine() {
	fileName := "./docs/data-20190514T0100.json"
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	var okvedItems []OkvedItem
	err = json.Unmarshal(data, &okvedItems)

	if err != nil {
		panic(err)
	}

	var sumGlobalId int64

	for _, item := range okvedItems {
		sumGlobalId += item.GlobalId
	}

	fmt.Printf("sumGlobalId: %d", sumGlobalId)
}

type OkvedItem struct {
	Idx            string `json:"Idx"`
	Kod            string `json:"Kod"`
	Name           string `json:"Name"`
	Razdel         string `json:"Razdel"`
	GlobalId       int64  `json:"global_id"`
	SignatureDate  string `json:"signature_date"`
	SystemObjectId string `json:"system_object_id"`
}
