package main

import (
	"fmt"
)

func SetBit(n int64, i uint, bit int) int64 {
	var mask int64 = 1 << i

	if bit == 1 {
		return n | mask
	} else if bit == 0 {
		return n &^ mask
	}
	return n
}

func main() {
	var number int64 = 5

	i1 := uint(1)
	bit1 := 0
	result1 := SetBit(number, i1, bit1)

	fmt.Printf("Число: %d (%b)\n", number, number)
	fmt.Printf("Сброс %d-го бита в 0:\n", i1)
	fmt.Printf("Результат: %d (%b)\n\n", result1, result1)

	number2 := int64(5)
	i2 := uint(2)
	bit2 := 1
	result2 := SetBit(number2, i2, bit2)

	fmt.Printf("Число: %d (%b)\n", number2, number2)
	fmt.Printf("Установка %d-го бита в 1:\n", i2)
	fmt.Printf("Результат: %d (%b)\n\n", result2, result2)

	number3 := int64(100)
	i3 := uint(6)
	bit3 := 0
	result3 := SetBit(number3, i3, bit3)

	fmt.Printf("Число: %d (%b)\n", number3, number3)
	fmt.Printf("Сброс %d-го бита в 0:\n", i3)
	fmt.Printf("Результат: %d (%b)\n", result3, result3)
}
