package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var firstName, lastName, surName, groupNumber string
	fmt.Print("Введите ваше имя: ")
	fmt.Scanln(&firstName)
	fmt.Print("Введите вашу фамилию:")
	fmt.Scanln(&lastName)
	fmt.Print("Введите ваше отчетсво: ")
	fmt.Scanln(&surName)
	fmt.Print("Введите номер группы: ")
	fmt.Scanln(&groupNumber)

	job := "* Лабораторная работа № 1"
	author := "* Выполнил(а): ст. гр. " + groupNumber
	fullName := "* " + lastName + " " + firstName + " " + surName

	var maxLen int
	if len(job) > len(author) && len(job) > len(fullName) {
		maxLen = utf8.RuneCountInString(job)

	} else if len(author) > len(fullName) && len(author) > len(job) {
		maxLen = utf8.RuneCountInString(author)
	} else {
		maxLen = utf8.RuneCountInString(fullName)
	}
	b := maxLen - utf8.RuneCountInString(job)
	for i := 0; i < b; i++ {
		job = job + " "
	}
	job += "*"
	b = maxLen - utf8.RuneCountInString(author)
	for i := 0; i < b; i++ {
		author = author + " "
	}
	author += "*"

	b = maxLen - utf8.RuneCountInString(fullName)
	for i := 0; i < b; i++ {
		fullName = fullName + " "
	}
	fullName += "*"

	a := ""
	for i := 0; i < maxLen+1; i++ {
		a = a + "*"
	}
	fmt.Println(a)

	fmt.Printf("%s\n%s\n%s\n", job, author, fullName)

	fmt.Println(a)

}

// * ddfudfushfudhfd f ifwofhoeifhwoihfeihfw  *
