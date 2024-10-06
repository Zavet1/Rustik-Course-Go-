package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Print("Введите первое число:")
	fmt.Scan(&a)
	fmt.Print("Введите второе число:")
	fmt.Scan(&b)
	fmt.Print("Введите третье число:")
	fmt.Scan(&c)

	if a == b || b == c || c == a {
		a += 5
		b += 5
		c += 5
		fmt.Printf("увеличение всех трех: %s, %s, %s\n", a, b, c)
	} else {
		fmt.Print("Равных нет")
	}
}
