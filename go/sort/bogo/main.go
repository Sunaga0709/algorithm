package main

import (
	"fmt"
	"math/rand"
)

func main() {
	bogoSort([]int{5, 6, 4, 9, 13, 39})
}

func bogoSort(numbers []int) []int {
	const limit = 1_000
	var cnt int
	for {
		cnt++

		if limit < cnt {
			fmt.Println("exceeded ...")
			break
		}

		rand.Shuffle(len(numbers), func(i, j int) {
			numbers[i], numbers[j] = numbers[j], numbers[i]
		})

		fmt.Printf("%3d ... %v\n", cnt, numbers)

		if inOrder(numbers) {
			break
		}
	}

	return numbers
}

func inOrder(numbers []int) bool {
	for i := range len(numbers) - 1 {
		if numbers[i] > numbers[i+1] {
			return false
		}
	}

	return true
}
