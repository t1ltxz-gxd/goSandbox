package tasks

import (
	"fmt"
	"time"
)

func Fifth() {
	// Количество секунд работы программы
	seconds := 5

	//читаем канал для передачи знач
	ch := make(chan int)

	// Запускаем горутину которая отправляет значения в канал
	go sendData(ch)

	// Читаем значения из канала
	for {
		select {
		case data := <-ch:
			fmt.Println("Received:", data)
		case <-time.After(time.Duration(seconds) * time.Second):
			fmt.Println("Timeout, program stopped.")
			return
		}
	}
}

func sendData(ch chan<- int) {
	for i := 1; ; i++ {
		time.Sleep(time.Second) // Задержка между отправкой значений в канал
		ch <- i
	}
}
