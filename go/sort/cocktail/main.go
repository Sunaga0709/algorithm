package main

import "fmt"

func main() {
	cocktailSort([]int{5, 7, 1, 34, 22, 59, 18, 9, 2})
}

func cocktailSort(numbers []int) {
	length := len(numbers)
	swap := true
	var start int
	end := length - 1
	var cnt int
	for {
		swap = false
		fmt.Println("~ backward ~")
		for i := start; i < end; i++ {
			if numbers[i] > numbers[i+1] {
				numbers[i], numbers[i+1] = numbers[i+1], numbers[i]
				swap = true
			}
			fmt.Printf("%3d_%d_%d ... %v\n", cnt, start, end, numbers)
		}

		if !swap {
			break
		}

		swap = false
		end--

		fmt.Println("~ forward ~")
		for i := end - 1; i > start-1; i-- {
			if numbers[i] > numbers[i+1] {
				numbers[i], numbers[i+1] = numbers[i+1], numbers[i]
				swap = true
			}
			fmt.Printf("%3d_%d_%d ... %v\n", cnt, start, end, numbers)
		}

		start++
		cnt++
	}
}
