package main

import "fmt"

func main() {
	fmt.Println("Вы заходите в кал поликлинику и видите огромную очередь")

	var people int64

	fmt.Print("Введите кол-во людей в очереди:")
	fmt.Scan(&people)

	minuets := people * 10
	hours := minuets / 60
	mid := minuets % 60

	fmt.Printf("вы туда сюда %d часов %d минут", hours, mid)

}
