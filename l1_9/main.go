package main

import (
	"fmt"
	"sync"
)

func generator(input []int, out chan<- int) {
	defer close(out)

	for _, x := range input {
		out <- x
	}
}

func processor(in <-chan int, out chan<- int) {
	defer close(out)

	for x := range in {
		result := x * 2
		out <- result
	}
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	firstChannel := make(chan int)
	secondChannel := make(chan int)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		generator(numbers, firstChannel)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		processor(firstChannel, secondChannel)
	}()

	fmt.Println("Результаты конвейера (x * 2):")

	for result := range secondChannel {
		fmt.Printf("%d\n", result)
	}

	wg.Wait()

	fmt.Println("Конвейер успешно завершен.")
}
