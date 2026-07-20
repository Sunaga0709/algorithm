package randomvalue

import "math/rand"

const maxValue = 999

func RandomSlice(length int) []int {
	value := make([]int, 0, length)
	for range length {
		val := rand.Intn(maxValue)
		value = append(value, val)
	}

	return value
}
