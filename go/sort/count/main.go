package main

import "fmt"

func main() {
	countSort([]int{5, 7, 1, 4, 2, 9, 8, 9, 2})
}

func countSort(numbers []int) {
	var max int
	for _, v := range numbers {
		if max < v {
			max = v
		}
	}
	fmt.Printf("  max: %d\n", max)
	fmt.Println("---")

	bucket := make([]int, max+1)
	for _, v := range numbers {
		bucket[v] = bucket[v] + 1
		fmt.Printf("  v: %d, ", v)
		fmt.Printf("bucket: %+v, ", bucket)
		fmt.Printf("buckert[v]: %d\n", bucket[v])
	}

	var result []int
	for i, v := range bucket {
		if v > 0 {
			for range v {
				result = append(result, i)
			}
		}
	}
	fmt.Println("---")

	fmt.Printf("  result: %v\n", result)
}
