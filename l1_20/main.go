package main

import "fmt"

func reverseRunesInPlace(runes []rune, i, j int) {
	for i < j {
		runes[i], runes[j] = runes[j], runes[i]
		i++
		j--
	}
}

func ReverseWords(s string) string {
	runes := []rune(s)
	n := len(runes)

	reverseRunesInPlace(runes, 0, n-1)

	start := 0
	for i := 0; i < n; i++ {
		if runes[i] == ' ' {
			reverseRunesInPlace(runes, start, i-1)
			start = i + 1
		}
	}

	reverseRunesInPlace(runes, start, n-1)

	return string(runes)
}

func main() {
	input := "snow dog sun"
	reversed := ReverseWords(input)
	fmt.Printf("Вход: \"%s\"\nВыход: \"%s\"\n", input, reversed)

	input2 := "hello wb tech school"
	reversed2 := ReverseWords(input2)
	fmt.Printf("Вход: \"%s\"\nВыход: \"%s\"\n", input2, reversed2)
}
