package main

import "fmt"

func RemoveElement(s []int, i int) []int {
	if i < 0 || i >= len(s) {
		return s
	}
	copy(s[i:], s[i+1:])
	s[len(s)-1] = 0
	return s[:len(s)-1]
}

func main() {
	data := []int{10, 20, 30, 40, 50}
	indexToRemove := 2

	fmt.Printf("Исходный срез: %v, Длина: %d, Емкость: %d\n", data, len(data), cap(data))

	result := RemoveElement(data, indexToRemove)

	fmt.Printf("Результат: %v, Длина: %d, Емкость: %d\n", result, len(result), cap(result))
	fmt.Printf("Исходный массив: %v\n", data)
}
