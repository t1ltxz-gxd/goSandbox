package tasks

import "fmt"

// Target - Целевой интерфейс, который адаптер должен реализовать.
type Target interface {
	Request() string
}

// Adaptee - Адаптируемый класс.
type Adaptee struct{}

// SpecificRequest - Метод адаптируемого класса.
func (a *Adaptee) SpecificRequest() string {
	return "Specific request"
}

// Adapter - Адаптер, реализующий целевой интерфейс и использующий адаптируемый класс.
type Adapter struct {
	adaptee *Adaptee
}

// Request - Метод адаптера, который вызывает метод адаптируемого класса.
func (a *Adapter) Request() string {
	return a.adaptee.SpecificRequest()
}

func TwentyFirst() {
	adaptee := &Adaptee{}
	adapter := &Adapter{
		adaptee: adaptee,
	}

	fmt.Println(adapter.Request())
}
