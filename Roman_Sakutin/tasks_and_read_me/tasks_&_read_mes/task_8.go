package main

import (
	"fmt"
)

func main() {
	pasword := "qwerty123"
	var inputPassword string
	maxAttempts := 5

	for attempts := 0; attempts < maxAttempts; attempts++ {
		fmt.Print("Введите пароль: ")
		fmt.Scan(&inputPassword)

		if inputPassword == pasword {
			fmt.Println("Доступ к тайному сообщению разрешен!")
			fmt.Println("Секретное сообщение: Вы успешно ввели пароль: good.")
			break
		} else {
			fmt.Println("Неверный пароль. Попробуйте еще раз.")
		}
	}

	// Если пользователь ввел неверный пароль 3 раза
	fmt.Println("Вы исчерпали все попытки. Программа завершена.")
}
