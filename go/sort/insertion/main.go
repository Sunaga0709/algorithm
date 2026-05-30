package main

import "fmt"

func main() {
	insertionSort([]int{5, 7, 1, 4, 2, 9, 8, 9, 2})
}

func insertionSort(numbers []int) {
	length := len(numbers)
	for i := range length {
		if i == 0 {
			continue
		}

		temp := numbers[i]
		fmt.Printf("   %2d: temp: %d\n", i, temp)
		j := i - 1
		for j >= 0 && numbers[j] > temp {
			numbers[j+1] = numbers[j]
			j--
			fmt.Printf("%2d_%2d: %+v\n", i, j, numbers)
		}
		numbers[j+1] = temp

		fmt.Printf("   %2d: %+v\n", i, numbers)
		fmt.Println("---")
	}

	fmt.Println("---")
	fmt.Printf("result: %+v\n", numbers)
}
