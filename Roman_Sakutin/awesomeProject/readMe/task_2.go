package main

import (
	"fmt"
)

func main() {
	var a string

	// Цикл выполняется, пока не введено слово "exit"
	for {
		fmt.Print("Введите команду (или 'exit' для выхода): ")
		fmt.Scanln(&a)

		if a == "exit" {
			fmt.Println("Выход из программы.")
			break
		}

		fmt.Println("Вы ввели:", a)
	}
}
