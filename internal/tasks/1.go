package tasks

import "fmt"

// Human Структура Human
type Human struct {
	name string
	age  int
}

// GetName Метод GetName для структуры Human
func (h Human) GetName() string {
	return h.name
}

// GetAge Метод GetAge для структуры Human
func (h Human) GetAge() int {
	return h.age
}

// Action Структура Action с встраиванием структуры Human
type Action struct {
	Human
}

// Walk Метод Walk для структуры Action
func (a Action) Walk() {
	fmt.Println(a.GetName(), "is walking.")
}

// Run Метод Run для структуры Action
func (a Action) Run() {
	fmt.Println(a.GetName(), "is running.")
}

func First() {
	// Создание объекта типа Action
	action := Action{
		Human: Human{
			name: "John",
			age:  25,
		},
	}

	// Вызов встроенных методов
	action.Walk()
	action.Run()
	fmt.Println("Name:", action.GetName())
	fmt.Println("Age:", action.GetAge())
}
