package main

import "fmt"

func ReverseString(s string) string {
	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func main() {
	input1 := "главрыба"
	reversed1 := ReverseString(input1)
	fmt.Printf("Вход: %s\nВыход: %s\n", input1, reversed1)

	input2 := "Hello, WB!"
	reversed2 := ReverseString(input2)
	fmt.Printf("Вход: %s\nВыход: %s\n", input2, reversed2)

	input3 := "abcde"
	reversed3 := ReverseString(input3)
	fmt.Printf("Вход: %s\nВыход: %s\n", input3, reversed3)
}
