package main

import "fmt"

func main() {
	bubbleSort([]int{5, 7, 1, 34, 22, 59, 18, 9, 2})
}

func bubbleSort(numbers []int) {
	length := len(numbers)
	for i := range length {
		// `i`が終端
		for j := range length - 1 - i {
			if numbers[j] > numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}

			fmt.Printf("%d_%d ... %v\n", i, j, numbers)
		}
		fmt.Println("---")
	}
}
