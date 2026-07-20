package main

import (
	"fmt"

	"github.com/Sunaga0709/algorithm/go/randomvalue"
)

func main() {
	mergeSort(randomvalue.RandomSlice(10))
}

func mergeSort(numbers []int) {
	fmt.Printf("input: %+v\n", numbers)

	if len(numbers) <= 1 {
		return
	}

	center := len(numbers) / 2
	left := append([]int(nil), numbers[:center]...)
	right := append([]int(nil), numbers[center:]...)

	mergeSort(left)
	mergeSort(right)

	var i, j, k int
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			numbers[k] = left[i]
			i += 1
		} else {
			numbers[k] = right[j]
			j += 1
		}

		k += 1
	}

	for i < len(left) {
		numbers[k] = left[i]
		i += 1
		k += 1
	}

	for j < len(right) {
		numbers[k] = right[j]
		j += 1
		k += 1
	}

	fmt.Printf("result: %+v\n", numbers)
}
