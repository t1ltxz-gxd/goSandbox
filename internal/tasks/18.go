package tasks

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu    sync.Mutex
	value int
	wg    sync.WaitGroup
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
	c.wg.Done()
}

func Eighteenth() {
	counter := Counter{value: 0}

	// Устанавливаем количество горутин
	numThreads := 10
	counter.wg.Add(numThreads)

	for i := 0; i < numThreads; i++ {
		go counter.Increment()
	}

	// Ждем завершения всех горутин
	counter.wg.Wait()

	// Выводим итоговое значение счетчика
	fmt.Println("Итоговое значение счетчика:", counter.value)
}
