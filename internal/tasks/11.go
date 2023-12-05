package tasks

import "fmt"

func intersection(set1, set2 []int) []int {
	intersectSet := make([]int, 0)
	setMap := make(map[int]bool)

	// Заполняем карту значениями из первого множества
	for _, num := range set1 {
		setMap[num] = true
	}

	// Проверяем значения второго множества и добавляем их в пересечение, если они есть в первом множестве
	for _, num := range set2 {
		if setMap[num] {
			intersectSet = append(intersectSet, num)
		}
	}

	return intersectSet
}

func Eleventh() {
	set1 := []int{1, 2, 3, 4, 5}
	set2 := []int{4, 5, 6, 7, 8}

	intersect := intersection(set1, set2)
	fmt.Println("Пересечение множеств:", intersect)
}
