package main

import "fmt"

func main() {
	shellSort([]int{51, 47, 11, 34, 102, 92, 84, 29, 52, 43})
}

func shellSort(numbers []int) {
	length := len(numbers)
	gap := length / 2

	for gap > 0 {
		fmt.Printf("--- gap: %d ---\n", gap)
		for i := gap; i < length; i++ {
			temp := numbers[i]
			j := i
			for j >= gap && numbers[j-gap] > temp {
				numbers[j] = numbers[j-gap]
				j -= gap
			}

			numbers[j] = temp
			fmt.Printf("temp: %3d, numbers: %+v\n", temp, numbers)
		}

		gap /= 2
	}

	fmt.Printf("---\nnumbers: %+v\n", numbers)
}
