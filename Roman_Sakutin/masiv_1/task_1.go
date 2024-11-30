package main

import "fmt"

func main() {
	matrix := [][]int{
		{4, 5, 6},    //0
		{7, 8, 9},    //1
		{10, 11, 12}, //2
	}
	sum_2 := 0
	for _, sum := range matrix[1] {
		sum_2 += sum
	}
	sum_1 := 1
	for _, suma := range matrix {
		sum_1 *= suma[0]
	}
	fmt.Println("Исходная матрица:")
	for _, c := range matrix {
		fmt.Println(c)
	}
	fmt.Printf("Сумма 2 строки: %d\n", sum_2)
	fmt.Printf("Произведение 1 столбца: %d\n", sum_1)
}
