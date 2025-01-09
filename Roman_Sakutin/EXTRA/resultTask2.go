package main

import (
	"fmt"
)

func main() {
	var a, b, f float32

	fmt.Println("Введите значения для a, b и f:")

	fmt.Print("a: ")
	fmt.Scan(&a)
	fmt.Print("b: ")
	fmt.Scan(&b)
	fmt.Print("f: ")
	fmt.Scan(&f)

	result := (a + b - f/a) + f*a*a - (a + b)

	fmt.Printf("Результат выражения: %f\n", result)
}
