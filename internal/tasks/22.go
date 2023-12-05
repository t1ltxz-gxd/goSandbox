package tasks

import "fmt"

func TwentySecond() {
	a := 2 << 20 // Значение переменной a, которое больше чем 2^20
	b := 3 << 20 // Значение переменной b, которое больше чем 2^20

	// Перемножение
	multiplication := a * b
	fmt.Println("Перемножение:", multiplication)

	// Деление
	division := a / b
	fmt.Println("Деление:", division)

	// Сложение
	addition := a + b
	fmt.Println("Сложение:", addition)

	// Вычитание
	subtraction := a - b
	fmt.Println("Вычитание:", subtraction)
}
