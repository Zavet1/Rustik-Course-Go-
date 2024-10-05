package main

import "fmt"

func main() {
	var firstName, lastName, surName, groupNumber string
	fmt.Print("Введите ваше имя: ")
	fmt.Scan(&firstName)
	fmt.Print("Введите вашу фамилию:")
	fmt.Scan(&lastName)
	fmt.Print("Введите ваше отчетсво: ")
	fmt.Scan(&surName)
	fmt.Print("Введите номер группы: ")
	fmt.Scan(&groupNumber)

	job := "Лабораторная работа № 1"
	author := "Выполнил(а): ст. гр. " + groupNumber
	fullName := lastName + " " + firstName + " " + surName
	fmt.Print(job + author + fullName)
}
