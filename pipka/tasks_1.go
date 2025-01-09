package main

import (
	"fmt"
)

func isValidAndMaxDepth(s string) (bool, int) {
	maxDepth := 0
	balance := 0

	for _, char := range s {
		if char == '(' {
			balance++
			if balance > maxDepth {
				maxDepth = balance
			}
		} else if char == ')' {
			balance--
			if balance < 0 {
				return false, 0 // Некорректная строка
			}
		}
	}

	return balance == 0, maxDepth
}

func main() {
	examples := []string{"(()(()))", "(()", "())", ")(", "(()))(()"}

	for _, example := range examples {
		valid, maxDepth := isValidAndMaxDepth(example)
		if valid {
			fmt.Printf("Строка \"%s\" корректная, максимальная глубина: %d\n", example, maxDepth)
		} else {
			fmt.Printf("Строка \"%s\" некорректная\n", example)
		}
	}
}
