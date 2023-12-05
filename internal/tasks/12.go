package tasks

import "fmt"

func Twelve() {
	sequence := []string{"cat", "cat", "dog", "cat", "tree"}

	uniqueSet := make(map[string]bool)
	for _, item := range sequence {
		uniqueSet[item] = true
	}

	fmt.Println("Уникальный набор строк:")
	for item := range uniqueSet {
		fmt.Println(item)
	}
}
