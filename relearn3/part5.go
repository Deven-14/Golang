package main

import (
	"fmt"
	"strings"
)

// * value or pointer receiver for methods

// * type parameters for generics

func Index[S ~[]E, E comparable](s S, v E) int {
	for i := range s {
		if v == s[i] {
			return i
		}
	}
	return -1
}

// * generic types
type List[T any] struct {
	next *List[T]
	val  T
}

func (l *List[T]) String() string {
	node := l
	values := []string{}
	for node != nil {
		values = append(values, fmt.Sprintf("%v", node.val))
		node = node.next
	}
	return strings.Join(values, ", ")
}

func part5() {

	// type parameters
	s := []int{10, 20, 30, 40}
	fmt.Println(Index(s, 15))
	fmt.Println(Index(s, 30))

	// generic types
	head := &List[int]{nil, 5}
	head.next = &List[int]{nil, 10}
	head.next.next = &List[int]{nil, 15}
	fmt.Println(head)

}
