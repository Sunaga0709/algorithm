package main

import (
	"fmt"
)

type node struct {
	data any
	next *node
}

func newNode(data any, next *node) node {
	return node{
		data: data,
		next: next,
	}
}

type linkedList struct {
	head *node
}

func newLinkedList(head *node) linkedList {
	return linkedList{
		head: head,
	}
}

func (l *linkedList) append(data any) {
	nd := newNode(data, nil)
	if l.head == nil {
		l.head = &nd
		return
	}

	lastNode := l.head
	for lastNode.next != nil {
		lastNode = lastNode.next
	}
	lastNode.next = &nd
}

func (l *linkedList) insert(data any) {
	nd := newNode(data, nil)
	nd.next = l.head
	l.head = &nd
}

func (l *linkedList) String() string {
	var result string
	if l.head == nil {
		result += "nil"
		return result
	}

	result += fmt.Sprintf("%+v", l.head.data)
	lastNode := l.head
	for lastNode.next != nil {
		lastNode = lastNode.next
		result += fmt.Sprintf(" -> %+v", lastNode.data)
	}

	return result
}

func main() {
	l := newLinkedList(nil)
	fmt.Printf("initialized: %#v , call String method: %s\n", l, l.String())
	l.append(1)
	fmt.Printf("appended 1: %#v\n", l)
	l.append("string")
	fmt.Printf("appended `string`: %#v\n", l)
	l.insert(true)
	fmt.Printf("inserted `true`: %#v\n", l)
	l.append(false)
	fmt.Printf("appended `false`: %#v\n", l)
	l.insert(1.23)
	fmt.Printf("inserted 1.23: %#v\n", l)

	fmt.Printf("------\ndata: %s\n", l.String())
}
