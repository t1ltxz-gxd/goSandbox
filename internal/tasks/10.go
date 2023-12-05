package tasks

import (
	"fmt"
	"math"
	"sort"
)

func groupTemperatures(temperatures []float64, step float64) map[int][]float64 {
	groups := make(map[int][]float64)
	for _, temp := range temperatures {
		groupKey := int(math.Floor(temp/step)) * int(step)
		groups[groupKey] = append(groups[groupKey], temp)
	}
	return groups
}

func Tenth() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	step := 10.0

	groups := groupTemperatures(temperatures, step)

	// Сортировка ключей в порядке возрастания
	var keys []int
	for key := range groups {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	// Вывод групп температур
	for _, key := range keys {
		fmt.Printf("%d: %v\n", key, groups[key])
	}
}
