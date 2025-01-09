package main

import "fmt"

func main() {
	var input rune
	fmt.Print("Введите латинскую букву в нижнем регистре: ")
	fmt.Scanf("%c", &input)
	fmt.Printf("Буква в верхнем регистре: %c\n", input-'a'+'A')
}
