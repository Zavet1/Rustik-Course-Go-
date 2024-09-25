package main

import "fmt"

func main() {
	fmt.Println("Введите ваше имя -")
	var (
		name string
		age  int
	)
	fmt.Scan(name)
	fmt.Println("Введите ваш возраст")
	fmt.Scan(age)
	fmt.Printf("Ваше имя %v. Ваш возраст %v ", name, age)

}
