package main

import (
	"fmt"
)

func main() {

	sum := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	sumSecond := 0
	for _, b := range sum[1] {
		sumSecond += b
	}

	first := 1
	for i := 0; i < len(sum); i++ {
		first *= sum[i][0]
	}

	fmt.Println("Исходная матрица:")
	for _, a := range sum {
		fmt.Println(a)
	}

	// Вывод результатов
	fmt.Printf("Сумма второй строки: %d\n", sumSecond)
	fmt.Printf("Произведение первого столбца: %d\n", first)
}
