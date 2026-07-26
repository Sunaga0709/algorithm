package main

import (
	"fmt"
)

func main() {
	numbers := []int{0, 1, 5, 7, 9, 11, 15, 20, 24}

	ind := binarySearch(numbers, 33)
	fmt.Printf("target index: %2d\n", ind)
}

//nolint:unused
func _binarySearch(numbers []int, value int) int {
	left, right := 0, len(numbers)-1
	for left <= right {
		mid := (left + right) / 2
		if numbers[mid] == value {
			return mid
		} else if numbers[mid] < value {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

func binarySearch(numbers []int, value int) int {
	var bs func(numbers []int, value, left, right int) int
	bs = func(numbers []int, value, left, right int) int {
		if left > right {
			return -1
		}

		mid := (left + right) / 2
		if numbers[mid] == value {
			return mid
		} else if numbers[mid] < value {
			return bs(numbers, value, mid+1, right)
		} else {
			return bs(numbers, value, left, mid-1)
		}
	}

	return bs(numbers, value, 0, len(numbers)-1)
}

//nolint:unused
func linearSearch(numbers []int, value int) int {
	for i := range len(numbers) {
		if numbers[i] == value {
			return i
		}
	}

	return -1
}
