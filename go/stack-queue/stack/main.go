package main

import "fmt"

type stack struct {
	value []any
}

func newStack() stack {
	return stack{
		value: []any{},
	}
}

func (s *stack) append(data any) {
	s.value = append([]any{data}, s.value...)
}

func (s *stack) pop() (any, bool) {
	if len(s.value) == 0 {
		return nil, false
	}

	val := s.value[0]
	s.value = s.value[1:]

	return val, true
}

func main() {
	s := newStack()
	fmt.Printf("initialized: %+v\n", s)

	s.append(1)
	fmt.Printf("added 1: %+v\n", s)

	s.append("string")
	fmt.Printf("added `string`: %+v\n", s)

	popped1, ok := s.pop()
	fmt.Printf("popped 1: %+v: value -> %v, ok -> %v\n", s, popped1, ok)

	popped2, ok := s.pop()
	fmt.Printf("popped 2: %+v: value -> %v, ok -> %v\n", s, popped2, ok)

	popped3, ok := s.pop()
	fmt.Printf("popped 3: %+v: value -> %v, ok -> %v\n", s, popped3, ok)
}
