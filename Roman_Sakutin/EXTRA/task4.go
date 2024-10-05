package main

import "fmt"

func main() {
	var meters float64

	fmt.Print("Метры:")
	fmt.Scan(&meters)
	kilomiters := meters
	fmt.Print("Километры: ")
	fmt.Print(kilomiters / 1000.0)

}
