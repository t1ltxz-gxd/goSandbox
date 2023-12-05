package tasks

import "fmt"

func setBit(n int64, pos uint, bitValue int) int64 {
	mask := int64(1 << pos)
	if bitValue == 1 {
		n |= mask
	} else {
		n &^= mask
	}
	return n
}

func Eighth() {
	var num int64 = 10 // Исходное значение
	pos := uint(2)     // Позиция бита (от 0 до 63)
	bitValue := 1      // Новое значение бита (0 или 1)

	result := setBit(num, pos, bitValue)
	fmt.Printf("Исходное значение: %d\n", num)
	fmt.Printf("Позиция бита: %d\n", pos)
	fmt.Printf("Новое значение: %d\n", result)
}
