package main

import "fmt"

func main() {
	fmt.Println("Введите ваше имя -")
	var (
		name string
		age  int
		znak string
		car  string
	)
	fmt.Scan(&name)
	fmt.Println("Введите ваш возраст")
	fmt.Scan(&age)
	fmt.Println("Введите ваш знак зодиака")
	fmt.Scan(&znak)
	fmt.Println("Любимая машина?")
	fmt.Scan(&car)
	fmt.Printf("Ваше имя %v. Ваш возраст %v. Знак задиака %v. Любимая машина %v. ", name, age, znak, car)

}
