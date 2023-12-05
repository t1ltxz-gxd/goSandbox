package tasks

import (
	"fmt"
	"sync"
)

func Third() {
	nums := []int{2, 4, 6, 8, 10}

	// Создаем ожидание для дожидания окончания выполнения всех горутин
	var wg sync.WaitGroup
	// Количество горутин равно количеству чисел в последовательности
	wg.Add(len(nums))

	// Канал для передачи значений квадратов чисел
	results := make(chan int)

	for _, num := range nums {
		// Конкурентно вычисляем квадрат числа
		go func(n int) {
			defer wg.Done()

			square := n * n
			results <- square
		}(num)
	}

	// Закрываем канал после завершения всех горутин
	go func() {
		wg.Wait()
		close(results)
	}()

	// Суммируем квадраты чисел, полученные из канала
	sum := 0
	for num := range results {
		sum += num
	}

	// Выводим результат
	fmt.Println("Сумма квадратов чисел:", sum)
}
