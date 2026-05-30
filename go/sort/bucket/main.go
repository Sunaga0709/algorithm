package main

import (
	"fmt"
	"slices"
)

func main() {
	bucketSort([]int{51, 47, 11, 34, 102, 92, 84, 29, 52, 43})
}

func bucketSort(numbers []int) {
	length := len(numbers)
	if length == 0 {
		fmt.Printf("result: %+v\n", numbers)
		return
	}

	minNumber := slices.Min(numbers)
	maxNumber := slices.Max(numbers)
	bucketSize := (maxNumber-minNumber)/length + 1

	fmt.Printf("---\nlength: %d\nminNumber: %d\nmaxNumber: %d\nbucketSize: %d\n---\n", length, minNumber, maxNumber, bucketSize)

	buckets := make([][]int, length)
	for _, v := range numbers {
		i := (v - minNumber) / bucketSize
		buckets[i] = append(buckets[i], v)
	}

	fmt.Printf("%+v\n", buckets)

	for i := range length {
		fmt.Printf("before sort: %+v ___ ", buckets[i])
		insertionSort(buckets[i])
		fmt.Printf("after sort: %+v\n", buckets[i])
	}

	result := make([]int, 0, length)
	for i, v := range buckets {
		fmt.Printf("%2d ___ %+v\n", i, v)
		result = append(result, v...)
	}

	fmt.Printf("---\nresult: %+v\n", result)
}

func insertionSort(numbers []int) []int {
	length := len(numbers)
	for i := range length {
		if i == 0 {
			continue
		}

		temp := numbers[i]
		j := i - 1
		for j >= 0 && numbers[j] > temp {
			numbers[j+1] = numbers[j]
			j--
		}
		numbers[j+1] = temp
	}

	return numbers
}
