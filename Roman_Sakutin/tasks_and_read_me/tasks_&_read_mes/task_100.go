package main

import (
	"fmt"
	"strconv"
)

func getInput() int {
	var input string
	var number int
	var err error

	for {
		fmt.Print("Введите число: ")
		fmt.Scanln(&input) // Запрашиваем ввод у пользователя

		number, err = strconv.Atoi(input) // Пробуем конвертировать строку в int
		if err == nil {
			break // Если конвертация успешна, выходим из цикла
		}
		fmt.Println("Некорректный ввод, попробуйте снова.") // Сообщаем об ошибке
	}
	return number // Возвращаем корректное число
}

func main() {
	result := getInput()                       // Получаем число от пользователя
	fmt.Printf("Вы ввели число: %d\n", result) // Выводим число в консоль
}
