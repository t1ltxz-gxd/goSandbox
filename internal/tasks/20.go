package tasks

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	words := strings.Fields(s)
	reversedWords := make([]string, len(words))

	for i, word := range words {
		reversedWords[len(words)-1-i] = word
	}

	return strings.Join(reversedWords, " ")
}

func Twentieth() {
	str := "snow dog sun — sun dog snow"
	reversedStr := reverseWords(str)

	fmt.Println("Исходная строка:", str)
	fmt.Println("Перевернутые слова:", reversedStr)
}
