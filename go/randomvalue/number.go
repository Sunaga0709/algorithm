package randomvalue

import "math/rand"

func RandomNumber(max int) int {
	return rand.Intn(max)
}
