package tasks

import "fmt"

func Thirteenth() {
	a, b := 10, 5

	fmt.Println("До обмена:")
	fmt.Println("a =", a)
	fmt.Println("b =", b)

	a = a + b
	b = a - b
	a = a - b

	fmt.Println("После обмена:")
	fmt.Println("a =", a)
	fmt.Println("b =", b)
}
