package main

import "fmt"

func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1
}

type List[T any] struct {
	next *List[T]
	val  T
}

func typeParameters() {
	s1 := []int{10, 20, 15, -10}
	fmt.Println(Index(s1, 15))

	s2 := []string{"foo", "bar", "baz"}
	fmt.Println(Index(s2, "hello"))

	head := List[int]{nil, 5}
	head.next = &List[int]{nil, 10}
	head.next.next = &List[int]{nil, 15}

	node := &head
	for node != nil {
		fmt.Println(node.val)
		node = node.next
	}

}
