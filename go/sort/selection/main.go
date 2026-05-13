package main

import "fmt"

func main() {
	selectionSort([]int{5, 6, 4, 9, 13, 39, 1})
}

func selectionSort(numbers []int) {
	numbersLen := len(numbers)
	for i := range numbersLen {
		minIndex := i
		for j := i + 1; j < numbersLen; j++ {
			if numbers[minIndex] > numbers[j] {
				minIndex = j
			}
		}

		numbers[i], numbers[minIndex] = numbers[minIndex], numbers[i]

		fmt.Printf("%3d numbers: %+v\n", i, numbers)
	}
}
