package main

import "fmt"

func main() {
	sequence := []string{"cat", "cat", "dog", "cat", "tree"}

	set := make(map[string]bool)

	for _, word := range sequence {
		set[word] = true
	}

	uniqueWords := make([]string, 0, len(set))
	for word := range set {
		uniqueWords = append(uniqueWords, word)
	}

	fmt.Printf("Исходная последовательность: %v\n", sequence)
	fmt.Printf("Множество (набор уникальных слов): %v\n", uniqueWords)
}
