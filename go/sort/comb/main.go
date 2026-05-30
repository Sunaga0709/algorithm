package main

import "fmt"

func main() {
	combSort([]int{5, 6, 4, 9, 13, 39})
}

func combSort(numbers []int) {
	numbersLen := len(numbers)
	gap := numbersLen
	swapped := true

	var cnt int
	for gap > 1 || swapped {
		cnt++

		gap = max(int(float64(gap)/1.3), 1)

		swapped = false
		for i := 0; i < numbersLen-gap; i++ {
			if numbers[i] > numbers[i+gap] {
				numbers[i], numbers[i+gap] = numbers[i+gap], numbers[i]
				swapped = true
			}

			fmt.Printf("%3d_%d gap: %d, swapped: %v, numbers: %v\n", cnt, i, gap, swapped, numbers)
		}
	}
}
