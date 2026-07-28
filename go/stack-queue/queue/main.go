package main

import "fmt"

type queue struct {
	value []any
}

func newQueue() queue {
	return queue{
		value: []any{},
	}
}

func (q *queue) enqueue(data any) {
	q.value = append(q.value, data)
}

func (q *queue) dequeue() (any, bool) {
	if len(q.value) == 0 {
		return nil, false
	}

	value := q.value[0]
	q.value = q.value[1:]

	return value, true
}

func main() {
	q := newQueue()
	fmt.Printf("initialized: %+v\n", q)

	q.enqueue(1)
	fmt.Printf("enqueued 1: %+v\n", q)

	q.enqueue("string")
	fmt.Printf("enqueued `string`: %+v\n", q)

	dequeued1, ok := q.dequeue()
	fmt.Printf("dequeued 1: %+v: value -> %v, ok -> %v\n", q, dequeued1, ok)

	dequeued2, ok := q.dequeue()
	fmt.Printf("dequeued 2: %+v: value -> %v, ok -> %v\n", q, dequeued2, ok)

	dequeued3, ok := q.dequeue()
	fmt.Printf("dequeued 3: %+v: value -> %v, ok -> %v\n", q, dequeued3, ok)
}
