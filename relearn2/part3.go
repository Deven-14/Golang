package main

import "fmt"

// type parameters
func index[T comparable](s []T, t T) int {
	for i, v := range s {
		if v == t {
			return i
		}
	}
	return -1
}

func typeParameters() {

	s := []int{10, 20, 35, 1235, 205}
	fmt.Println(index(s, 35))

}

// generic types
type Node[T any] struct {
	val  T
	next *Node[T]
}

func genericTypes() {
	head := &Node[int]{val: 5}
	head.next = &Node[int]{val: 6}
	head.next.next = &Node[int]{val: 19}

	node := head
	for node != nil {
		fmt.Println(node.val)
		node = node.next
	}
}

func part3() {
	typeParameters()
	genericTypes()
}
