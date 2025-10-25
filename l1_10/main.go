package main

import (
	"fmt"
	"math"
	"sort"
)

func getGroupKey(temp float64) int {
	if temp >= 0 {
		return int(math.Floor(temp/10.0) * 10.0)
	}

	return int(math.Ceil(temp/10.0) * 10.0)
}

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	groupedTemps := make(map[int][]float64)

	for _, temp := range temperatures {
		key := getGroupKey(temp)
		groupedTemps[key] = append(groupedTemps[key], temp)
	}

	keys := make([]int, 0, len(groupedTemps))
	for key := range groupedTemps {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	fmt.Println("Результат:")
	for _, key := range keys {
		fmt.Printf("%d: %v\n", key, groupedTemps[key])
	}
}
