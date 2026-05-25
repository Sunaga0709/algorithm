package main

import "fmt"

func main() {
	gnomeSort([]int{5, 7, 1, 4, 2, 9, 8, 9, 2})
}

func gnomeSort(numbers []int) {
	var index int
	for index < len(numbers) {
		if index == 0 || numbers[index-1] <= numbers[index] {
			index += 1
		} else {
			numbers[index-1], numbers[index] = numbers[index], numbers[index-1]
			index -= 1
		}
		fmt.Printf("%2d: %+v\n", index, numbers)
	}

	fmt.Println("---")
	fmt.Printf("result: %+v\n", numbers)
}
