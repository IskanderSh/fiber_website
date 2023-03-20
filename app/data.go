package app

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// Чтение данных из JSON-файла с переводом в слайс, состоящий из структур Student
func readDataJSON(id string) []Student {
	// Чтение из JSON-файла и логирование в случае ошибки
	file, err := ioutil.ReadFile("./data.json")
	if err != nil {
		log.Fatal(err)
	}

	// Объявление слайса из структур Student
	var data []Student

	// Запись из JSON-формата (который был строкой) в слайс data и логирование в случае ошибки
	if err := json.Unmarshal([]byte(file), &data); err != nil {
		log.Fatal(err)
	}

	// id = "all", когда пользователю выводятся все данные
	if id == "all" {

		// Возвращение данных
		return data
	}

	// Объявление переменной item в случае, когда id принадлежит конкретному животному
	var item Student

	// Выбор значения для item по id
	for i := 0; i < len(data); i++ {
		if data[i].Id == id {
			item = data[i]
		}
	}

	// Возвращение данных в виде слайса, т.к. возвращаемый тип функции - слайс из структур Student
	return []Student{item}
}

// Запись данных в JSON-файл
func writeDataJSON(item Student) error {
	// Чтение всех данных из исходного файла с данными
	data := readDataJSON("all")

	// Добавление нового студента в слайс с данными
	data = append(data, item)

	// Преобразование слайса из структур Student в набор байт, возвращение ошибки в случае, если она есть
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// Запись данных в исходный файл data.json с правами доступа 0644, возвращение ошибки в случае, если она есть
	if err := ioutil.WriteFile("./data.json", dataBytes, 0644); err != nil {
		return err
	}

	// Возвращение ошибки в случае, если она есть
	return err
}

// Обновление данных у одного студента
func updateDataJSON(item Student) error {
	// Чтение всех данных из исходного файла с данными
	data := readDataJSON("all")

	// Нахождение прежних данных о студенте по id и запись новых из item
	for i := 0; i < len(data); i++ {
		if data[i].Id == item.Id {
			data[i] = item
		}
	}

	// Преобразование слайса из структур Student в набор байт, возвращение ошибки в случае, если она есть
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// Запись данных в исходный файл data.json с правами доступа 0644, возвращение ошибки в случае, если она есть
	if err := ioutil.WriteFile("./data.json", dataBytes, 0644); err != nil {
		return err
	}

	// Возвращение ошибки в случае, если она есть
	return err
}

// Удаление студента по id из полученного слайса
func removeStudentFromSlice(s []Student, id string) []Student {
	i := 0

	// Поиск численного индекса из слайса по id
	for index, item := range s {
		if item.Id == id {
			i = index
		}
	}
	// Возвращение нового слайса без элемента с переданным id
	return append(s[:i], s[i+1:]...)
}
