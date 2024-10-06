package main

import "fmt"

func main() {
	var num1, num2 float64
	fmt.Println("Введите 1 число:")
	fmt.Scanln(&num1)
	fmt.Println("Введите 2 число")
	fmt.Scan(&num2)

	if num1 > num2 {
		fmt.Print("Первое число больше")
	} else if num1 < num2 {
		fmt.Print("Первое число меньше ")
	} else if num1 == num2 {
		fmt.Print("Эти числа равны")
	}

}
