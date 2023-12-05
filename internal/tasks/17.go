package tasks

import (
	"fmt"
	"sort"
)

func Seventeenth() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	target := 6

	index := sort.SearchInts(arr, target)

	if index < len(arr) && arr[index] == target {
		fmt.Printf("Элемент %d найден на позиции %d\n", target, index)
	} else {
		fmt.Printf("Элемент %d не найден в массиве\n", target)
	}
}
