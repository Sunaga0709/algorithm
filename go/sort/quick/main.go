package main

import (
	"fmt"

	"github.com/Sunaga0709/algorithm/go/randomvalue"
)

func main() {
	// quickSort([]int{51, 47, 11, 34, 102, 92, 84, 29, 52, 43})
	quickSort(randomvalue.RandomSlice(10))
}

func quickSort(numbers []int) {
	fmt.Printf("first numbers: %+v\n", numbers)
	var qs func(numbers []int, low, high int)
	qs = func(numbers []int, low, high int) {
		if low < high {
			partitionIndex := partition(numbers, low, high)
			fmt.Printf("partitionIndex: %2d, partitionValue: %3d, numbers: %+v\n", partitionIndex, numbers[partitionIndex], numbers)
			qs(numbers, low, partitionIndex-1)
			qs(numbers, partitionIndex+1, high)
		}
	}

	qs(numbers, 0, len(numbers)-1)

	fmt.Printf("result: %+v\n", numbers)
}

func partition(numbers []int, low, high int) int {
	i := low - 1
	pivot := numbers[high]

	for j := low; j < high; j++ {
		if numbers[j] <= pivot {
			i += 1
			numbers[i], numbers[j] = numbers[j], numbers[i]
		}
	}

	numbers[i+1], numbers[high] = numbers[high], numbers[i+1]

	return i + 1
}
