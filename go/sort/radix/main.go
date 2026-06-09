package main

import (
	"fmt"
	"slices"
)

func main() {
	radixSort([]int{125, 87, 41, 444, 952, 19, 8, 39, 2})
}

func radixSort(numbers []int) {
	maxNumber := slices.Max(numbers)
	place := 1

	for maxNumber > place {
		numbers = countSort(numbers, place)
		fmt.Printf("place: %5d: %+v\n", place, numbers)
		place *= 10
	}

	fmt.Printf("result: %+v\n", numbers)
}

func countSort(numbers []int, place int) []int {
	counts := make([]int, 10)
	result := make([]int, len(numbers))

	for _, n := range numbers {
		index := int(n/place) % 10
		counts[index] = counts[index] + 1
	}

	for i := 1; i < 10; i++ {
		counts[i] += counts[i-1]
	}

	i := len(numbers) - 1
	for i >= 0 {
		index := int(numbers[i]/place) % 10
		result[counts[index]-1] = numbers[i]
		counts[index] = counts[index] - 1
		i--
	}

	return result
}
