package tasks

import (
	"fmt"
	"sort"
)

func Sixteenth() {
	arr := []int{9, 3, 8, 5, 1, 4, 7, 6, 2}

	fmt.Println("Иходный массив:", arr)

	sort.Ints(arr)

	fmt.Println("Отсортированный массив:", arr)
}
