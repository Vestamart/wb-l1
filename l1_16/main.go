package main

import "fmt"

func quickSort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1
	pivotIndex := (left + right) / 2

	a[pivotIndex], a[right] = a[right], a[pivotIndex]
	i := left

	for j := left; j < right; j++ {
		if a[j] <= a[right] {
			a[i], a[j] = a[j], a[i]
			i++
		}
	}

	a[i], a[right] = a[right], a[i]

	quickSort(a[:i])
	quickSort(a[i+1:])

	return a
}

func main() {
	data1 := []int{10, 80, 30, 90, 40, 50, 70}
	fmt.Println("Исходный срез:", data1)
	sortedData1 := quickSort(data1)
	fmt.Println("Отсортированный срез:", sortedData1)

	data2 := []int{5, 2, 4, 6, 1, 3}
	fmt.Println("Исходный срез:", data2)
	sortedData2 := quickSort(data2)
	fmt.Println("Отсортированный срез:", sortedData2)
}
