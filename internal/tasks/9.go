package tasks

import (
	"fmt"
	"sync"
)

func Ninth() {
	in := make(chan int)
	out := make(chan int)

	// Запись чисел в первый канал
	go func() {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		for _, num := range numbers {
			in <- num
		}
		close(in)
	}()

	// Выполнение операции x*2 и запись результата во второй канал
	go func() {
		for num := range in {
			out <- num * 2
		}
		close(out)
	}()

	// Вывод результатов из второго канала в stdout
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		for num := range out {
			fmt.Println(num)
		}
	}()

	wg.Wait()
}
