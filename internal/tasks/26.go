package tasks

import (
	"fmt"
	"strings"
)

// IsUnique - функция, которая проверяет, являются ли все символы в строке уникальными (регистронезависимо)
func IsUnique(str string) bool {
	lowercaseStr := strings.ToLower(str)
	seen := make(map[rune]bool)

	for _, char := range lowercaseStr {
		if seen[char] {
			return false
		}
		seen[char] = true
	}

	return true
}

func TwentySixth() {
	str1 := "abcd"
	fmt.Println(str1, "-", IsUnique(str1))

	str2 := "abCdefAaf"
	fmt.Println(str2, "-", IsUnique(str2))

	str3 := "   aabcd"
	fmt.Println(str3, "-", IsUnique(str3))
}
