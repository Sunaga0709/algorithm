package main

import (
	"fmt"
	"slices"
)

func main() {
	countSort([]int{5, 7, 1, 4, 2, 9, 8, 9, 2})
}

func countSort(numbers []int) {
	maxNumber := slices.Max(numbers)
	counts := make([]int, maxNumber+1)
	result := make([]int, len(numbers))

	fmt.Println(numbers)
	for _, n := range numbers {
		counts[n] = counts[n] + 1
	}

	for i := 1; i < len(counts); i++ {
		counts[i] += counts[i-1]
	}

	i := len(numbers) - 1
	for i >= 0 {
		index := numbers[i]
		result[counts[index]-1] = numbers[i]
		counts[index] = counts[index] - 1
		i--
	}

	fmt.Printf("%+v\n", result)
}
