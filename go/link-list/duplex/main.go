package main

import "fmt"

type node struct {
	data any
	prev *node
	next *node
}

func newNode(data any, prev, next *node) node {
	return node{
		data: data,
		prev: prev,
		next: next,
	}
}

type duplexLinkedList struct {
	head *node
}

func newDuplexLinkedList(head *node) duplexLinkedList {
	return duplexLinkedList{
		head: head,
	}
}

func (d *duplexLinkedList) append(data any) {
	nd := newNode(data, nil, nil)

	if d.head == nil {
		d.head = &nd
		return
	}

	current := d.head
	for current.next != nil {
		current = current.next
	}
	current.next = &nd
	nd.prev = current
}

func (d *duplexLinkedList) insert(data any) {
	nd := newNode(data, nil, nil)

	if d.head == nil {
		d.head = &nd
		return
	}

	d.head.prev = &nd
	nd.next = d.head
	d.head = &nd
}

func (d *duplexLinkedList) String() string {
	var result string
	if d.head == nil {
		result += "nil"
		return result
	}

	result += fmt.Sprintf("%+v", d.head.data)
	lastNode := d.head
	for lastNode.next != nil {
		lastNode = lastNode.next
		result += fmt.Sprintf(" -> %+v", lastNode.data)
	}

	return result
}

func main() {
	d := newDuplexLinkedList(nil)
	fmt.Printf("initialized: %s\n", d.String())
	d.append(1)
	fmt.Printf("appended 1: %s\n", d.String())
	d.append(2)
	fmt.Printf("appended 2: %s\n", d.String())
	d.append("string")
	fmt.Printf("appended `string`: %s\n", d.String())
	d.insert(true)
	fmt.Printf("inserted `true`: %s\n", d.String())
	d.append(false)
	fmt.Printf("appended `false`: %s\n", d.String())

	fmt.Println("-----")
	fmt.Printf("head -> next:\n => %#v\n", d.head.next)
	fmt.Printf("head -> next -> next -> prev:\n => %#v\n", d.head.next.next.prev)
}
