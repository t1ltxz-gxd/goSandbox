package tasks

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func Fourth() {
	// Количество воркеров
	numWorkers := 5

	// Создаем канал для передачи данных
	dataCh := make(chan int)

	// Создаем ожидание для дожидания окончания выполнения всех воркеров
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	// Запускаем воркеры
	for i := 0; i < numWorkers; i++ {
		go worker(i, dataCh, &wg)
	}

	// Записываем данные в канал
	go func() {
		defer close(dataCh)

		for i := 0; ; i++ {
			dataCh <- i
		}
	}()

	// Создаем канал для получения сигнала об окончании программы по нажатию Ctrl+C
	stopCh := make(chan os.Signal, 1)
	signal.Notify(stopCh, os.Interrupt, syscall.SIGTERM)

	// Ожидаем получения сигнала об окончании программы
	<-stopCh

	// Останавливаем все воркеры
	fmt.Println("Stopping workers...")
	close(dataCh) // Закрываем канал данных, чтобы воркеры могли завершиться
	wg.Wait()     // Дожидаемся окончания выполнения всех воркеров
	fmt.Println("Program stopped.")
}

func worker(id int, dataCh <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Worker %d started\n", id)

	for data := range dataCh {
		fmt.Printf("Worker %d received: %d\n", id, data)
	}

	fmt.Printf("Worker %d stopped\n", id)
}
