package main

import "fmt"

func main() {
	var meters float64

	fmt.Print("Метры:")
	fmt.Scan(&meters)

	kilometres := meters / 1000.0

	fmt.Print(kilometres / 1000.0)
}
