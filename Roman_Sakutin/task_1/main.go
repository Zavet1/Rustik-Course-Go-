package main

import "fmt"

func main() {
	a := 1
	b := 2
	c := []int{1, 2, 3, 4, 5} // ну вроде как тут все должно быть гуд с интом
	d := 10 + 12
	e := 22 == d //я чета написал букву "e" а он ошибку выдавал в итоге "_" будет
	fmt.Println(c, a, b)
	fmt.Println(e)

}
