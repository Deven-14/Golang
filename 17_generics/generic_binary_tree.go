package main

import "fmt"

type Tree[T interface{}] struct {
	left, right *Tree[T]
	data        T
}

func (t *Tree[T]) Loookup(x T) *Tree[T] {
	return &Tree[T]{}
}

var stringTree Tree[string] = Tree[string]{
	data:  "omg",
	left:  nil,
	right: nil,
}

func Abc() {
	fmt.Println(stringTree)
}
