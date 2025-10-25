package main

import "fmt"

func Intersection(sliceA, sliceB []int) []int {
	setA := make(map[int]bool)
	for _, item := range sliceA {
		setA[item] = true
	}

	var intersection []int

	for _, item := range sliceB {
		if _, found := setA[item]; found {
			intersection = append(intersection, item)

			delete(setA, item)
		}
	}

	return intersection
}

func main() {
	A := []int{1, 2, 3}
	B := []int{2, 3, 4}

	result := Intersection(A, B)

	fmt.Printf("Множество A: %v\n", A)
	fmt.Printf("Множество B: %v\n", B)
	fmt.Printf("Пересечение: %v\n", result)
}
