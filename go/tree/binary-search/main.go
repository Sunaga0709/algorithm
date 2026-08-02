package main

import (
	"fmt"

	"github.com/Sunaga0709/algorithm/go/randomvalue"
)

type node struct {
	value int
	left  *node
	right *node
}

func (n *node) debugPrint() {
	n.debugPrintRotated("", "")
}

func (n *node) debugPrintRotated(prefix string, childrenPrefix string) {
	if n == nil {
		return
	}

	n.right.debugPrintRotated(childrenPrefix+"┌── ", childrenPrefix+"│   ")
	fmt.Printf("%s%d\n", prefix, n.value)
	n.left.debugPrintRotated(childrenPrefix+"└── ", childrenPrefix+"    ")
}

func newNode(value int) *node {
	return &node{
		value: value,
		left:  nil,
		right: nil,
	}
}

func insert(node *node, value int) *node {
	if node == nil {
		return newNode(value)
	}

	if value < node.value {
		node.left = insert(node.left, value)
	} else {
		node.right = insert(node.right, value)
	}

	return node
}

func inorder(node *node) {
	if node != nil {
		inorder(node.left)
		fmt.Println(node.value)
		inorder(node.right)
	}

}

func search(node *node, value int) bool {
	switch {
	case node == nil:
		return false
	case node.value == value:
		return true
	case node.value > value:
		return search(node.left, value)
	default:
		return search(node.right, value)
	}
}

func _minValue(node *node) *node {
	current := node
	for current.left != nil {
		current = current.left
	}

	return current
}

func remove(node *node, value int) *node {
	switch {
	case node == nil:
	case node.value == value:
		switch {
		case node.left == nil:
			return node.right
		case node.right == nil:
			return node.left
		default:
			tmp := _minValue(node)
			node.value = tmp.value
			node.right = remove(node.right, tmp.value)
		}
	case node.value > value:
		node.left = remove(node.left, value)
	default:
		node.right = remove(node.right, value)
	}

	return node
}

func main() {
	var root *node
	root = insert(root, 3)
	root = insert(root, 6)
	root = insert(root, 1)
	root = insert(root, 2)
	root = insert(root, 9)
	root = insert(root, 0)
	root = insert(root, 7)

	inorder(root)

	for range 3 {
		n := randomvalue.RandomNumber(10)
		fmt.Printf("search %2d: %v\n", n, search(root, n))
	}

	fmt.Println("===============")
	fmt.Println("before")
	root.debugPrint()
	fmt.Printf("----------\nafter\n")
	root = remove(root, 3)
	root.debugPrint()
}
