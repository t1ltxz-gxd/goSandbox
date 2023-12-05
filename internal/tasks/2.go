package tasks

import (
	"fmt"
	"sync"
)

func Second() {
	nums := []int{2, 4, 6, 8, 10}

	// Создаем ожидание для дожидания окончания выполнения всех горутин
	var wg sync.WaitGroup
	// Количество горутин равно количеству чисел в массиве
	wg.Add(len(nums))

	for _, num := range nums {
		// Конкурентно вычисляем квадрат числа
		go func(n int) {
			defer wg.Done()

			square := n * n
			// Выводим квадрат числа в stdout
			fmt.Println(square)
		}(num)
	}

	// Ожидаем окончания выполнения всех горутин
	wg.Wait()
}
