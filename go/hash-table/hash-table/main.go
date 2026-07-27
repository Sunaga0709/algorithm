package main

import (
	"crypto/md5"
	"fmt"
	"math/big"
)

const defaultIndexSize = 10

type hashTable struct {
	size  int
	table [][]*hashTableData
}

type hashTableData struct {
	key   string
	value any
}

func newHashTable(size *int) hashTable {
	s := defaultIndexSize
	if size != nil {
		s = *size
	}

	table := make([][]*hashTableData, s)

	return hashTable{
		size:  s,
		table: table,
	}
}

func (h *hashTable) hash(key string) int {
	sum := md5.Sum([]byte(key))
	return int(new(big.Int).Mod(new(big.Int).SetBytes(sum[:]), big.NewInt(int64(h.size))).Int64())
}

func (h *hashTable) add(key string, value any) {
	index := h.hash(key)

	for _, v := range h.table[index] {
		if v.key == key {
			v.value = value
			return
		}
	}

	h.table[index] = append(h.table[index], &hashTableData{key: key, value: value})
}

func (h *hashTable) get(key string) (any, bool) {
	index := h.hash(key)
	for _, v := range h.table[index] {
		if v.key == key {
			return v.value, true
		}
	}

	return nil, false
}

func (h *hashTable) debugPrint() {
	for idx, values := range h.table {
		fmt.Printf("index: %d, value: [", idx)
		for i, value := range values {
			if i > 0 {
				fmt.Print(" ")
			}
			fmt.Printf("%+v", *value)
		}
		fmt.Println("]")
	}
}

func main() {
	ht := newHashTable(nil)

	ht.add("car", "toyota")
	ht.add("flag", true)
	ht.add("age", 23)
	ht.add("example", "sample")

	ht.debugPrint()

	fmt.Println("--------------")

	car, ok := ht.get("car")
	fmt.Printf("get car: car -> %v, ok -> %v\n", car, ok)

	flag, ok := ht.get("flag")
	fmt.Printf("get flag: flag -> %v, ok -> %v\n", flag, ok)

	unknown, ok := ht.get("unknown")
	fmt.Printf("get unknown: unknown -> %v, ok -> %v\n", unknown, ok)
}
