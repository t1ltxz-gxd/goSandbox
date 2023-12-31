package tasks

import "context"

// Использование команды return или break:
//Вы можете использовать команду return или break внутри горутины для остановки ее выполнения. Это работает,
//если вы находитесь в цикле или функции, и хотите прекратить выполнение горутины после выполнения определенной
//логики.

//Использование канала:
//Вы можете использовать канал для организации взаимодействия между горутинами. Создайте канал и отправьте в него
//сигнал, чтобы указать горутине остановиться. Горутина будет читать значение из канала и выйдет из своего цикла или
//функции при получении определенного сигнала.

func Sixth() {
	stop := make(chan bool)
	go func() {
		for {
			select {
			case <-stop:
				return
			}
			// выполнять определенную логику
		}
	}()

	// Отправка сигнала горутине для остановки
	stop <- true

	/*
		====================================================================================================================
	*/

	// Использование контекста:
	//В Go 1.7 и выше появилось понятие контекста, который позволяет передавать сигналы для остановки выполнения
	//горутины. Вы можете использовать пакет context и функцию WithCancel для создания контекста и функцию Cancel для
	//остановки выполнения горутины.
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			}
		}
	}()

	// Отмена выполнения горутины
	cancel()

	/*
		====================================================================================================================
	*/

	// Использование флага для остановки выполнения:
	//Вы можете использовать флаг или булевую переменную, чтобы указать горутине остановиться. Горутина будет проверять
	//значение флага или переменной и выйдет из своего цикла или функции, если значение станет равным true.
	var stopTwo bool

	go func() {
		for {
			if stopTwo {
				return
			}
			// выполнять определенную логику
		}
	}()

	// Установка флага для остановки горутины
	stopTwo = true
}
