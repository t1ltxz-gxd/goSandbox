package tasks

import (
	"fmt"
	"unicode/utf8"
)

func reverseString(s string) string {
	var reversed string
	length := utf8.RuneCountInString(s)

	for i := length - 1; i >= 0; i-- {
		runeValue, _ := utf8.DecodeLastRuneInString(s[:i+1])
		reversed += string(runeValue)
	}

	return reversed
}

func Nineteenth() {
	str := "главрыба — абырвалг"
	reversedStr := reverseString(str)

	fmt.Println("Исходная строка:", str)
	fmt.Println("Перевернутая строка:", reversedStr)
}
