package main

import "fmt"

func binarySearch(arr []int, target int) int {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := low + (high-low)/2

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			low = mid + 1
		} else { // arr[mid] > target
			high = mid - 1
		}
	}

	return -1
}

func main() {
	sortedArray := []int{2, 5, 8, 12, 16, 23, 38, 56, 72, 91}

	target1 := 23
	index1 := binarySearch(sortedArray, target1)
	fmt.Printf("Поиск %d: Индекс = %d (Ожидается 5)\n", target1, index1)

	target2 := 42
	index2 := binarySearch(sortedArray, target2)
	fmt.Printf("Поиск %d: Индекс = %d (Ожидается -1)\n", target2, index2)

	target3 := 2
	index3 := binarySearch(sortedArray, target3)
	fmt.Printf("Поиск %d: Индекс = %d (Ожидается 0)\n", target3, index3)
}
