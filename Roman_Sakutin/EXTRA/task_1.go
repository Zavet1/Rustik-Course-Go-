package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Введите пятизначное число: ")

	var number int

	fmt.Scan(&number)

	numberString := strconv.Itoa(number)

	for i := len(numberString) - 1; i >= 0; i-- {
		fmt.Println(string(numberString[i]))

	}
}
