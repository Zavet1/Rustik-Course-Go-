package main

import "fmt"

func main() {
	var a, name string

	fmt.Print("Введите символ: ")
	fmt.Scan(&a)

	fmt.Print("Введите имя: ")
	fmt.Scan(&name)

	length := len(name) + 2

	for i := 0; i < length; i++ {
		fmt.Print(a)
	}
	fmt.Println()

	fmt.Printf("%v%v%v\n", a, name, a)

	for i := 0; i < length; i++ {
		fmt.Print(a)
	}
	fmt.Println()
}
