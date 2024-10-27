package main

import "fmt"

func main() {

	var a string
	var b int

	fmt.Print("Введите сообщение: ")
	fmt.Scanln(&a)
	fmt.Print("Введите количество повторов: ")
	fmt.Scanln(&b)

	for i := 0; i < b; i++ {
		fmt.Println(a)
	}
}
