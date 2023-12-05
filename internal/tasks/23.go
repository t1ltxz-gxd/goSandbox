package tasks

import "fmt"

func TwentyThree() {
	slice := []int{1, 2, 3, 4, 5}
	i := 2 //декс элемента, который нужно удалить

	// Удаление элемента из слайса
	slice = append(slice[:i], slice[i+1:]...)

	fmt.Println("Слайс после удаления i-ого элемента:", slice)
}
