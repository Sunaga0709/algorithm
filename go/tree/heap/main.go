package main

import (
	"fmt"
	"math"
)

type miniHeap struct {
	heap        []int
	currentSize int
}

func newMiniHeap() miniHeap {
	return miniHeap{
		heap:        []int{math.MinInt},
		currentSize: 0,
	}
}

func (m *miniHeap) parentIndex(index int) int {
	return index / 2
}

func (m *miniHeap) leftChildIndex(index int) int {
	return index * 2
}

func (m *miniHeap) rightChildIndex(index int) int {
	return (index * 2) + 1
}

func (m *miniHeap) swap(index1, index2 int) {
	m.heap[index1], m.heap[index2] = m.heap[index2], m.heap[index1]
}

func (m *miniHeap) heapifyUp(index int) {
	for m.parentIndex(index) > 0 {
		parentIndex := m.parentIndex(index)
		if m.heap[index] < m.heap[parentIndex] {
			m.swap(index, parentIndex)
		}

		index = parentIndex
	}
}

func (m *miniHeap) heapifyDown(index int) {
	for m.leftChildIndex(index) <= m.currentSize {
		miniChildIndex := m.miniChild(index)
		if m.heap[index] > m.heap[miniChildIndex] {
			m.swap(index, miniChildIndex)
		}

		index = miniChildIndex
	}
}

func (m *miniHeap) miniChild(index int) int {
	leftChildIndex := m.leftChildIndex(index)
	rightChildIndex := m.rightChildIndex(index)

	if rightChildIndex > m.currentSize {
		return leftChildIndex
	}

	if m.heap[leftChildIndex] < m.heap[rightChildIndex] {
		return leftChildIndex
	}

	return rightChildIndex
}

func (m *miniHeap) push(value int) {
	m.heap = append(m.heap, value)
	m.currentSize += 1
	m.heapifyUp(m.currentSize)
}

func (m *miniHeap) pop() (int, bool) {
	if len(m.heap) == 1 {
		return 0, false
	}

	root := m.heap[1]
	data := m.heap[len(m.heap)-1]
	m.heap = m.heap[:len(m.heap)-1]
	m.currentSize -= 1

	if len(m.heap) == 1 {
		return root, true
	}

	m.heap[1] = data
	m.heapifyDown(1)

	return root, true
}

func main() {
	m := newMiniHeap()
	m.push(5)
	m.push(6)
	m.push(2)
	m.push(9)
	m.push(13)
	m.push(11)
	m.push(1)
	fmt.Println(m.heap)
	m.pop()
	fmt.Println(m.heap)
}
