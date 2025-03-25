package main

import (
	"fmt"
	"sort"
)

// type parameter list
// [P, Q constraint1, R constraint2]

func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

// type constraint (they are meta-types, i.e., type for type V Number, here V is type and Number is type constraint (meta-type))
type Number interface {
	int64 | float64
}

func SumNumbers[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func main() {

	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	// fmt.Printf("Generic Sums: %v and %v\n",
	// 	SumIntsOrFloats[string, int64](ints),
	// 	SumIntsOrFloats[string, float64](floats))

	fmt.Printf("Generic Sums: %v and %v\n",
		SumIntsOrFloats(ints),
		SumIntsOrFloats(floats))

	fmt.Printf("Generic Sums Using type constraint: %v and %v\n",
		SumNumbers(ints),
		SumNumbers(floats))

}

// when to use generics

// wrtie go programs by writing code and not by defining types
// when it comes to generics if you start writing your program by defining new type parameter constraint you're probably on the wrong path, start by writing functions, it's easy to add type parameters later when it's clear that they'll be useful

//* when are type parameters useful
// 1) Functions that work on slices, maps, and channels of any element type.
// 2) General purpose data structures
// 2.1) When operating on type parameters, prefer functions to methods // i.e., adding cmp func coz if we constraint T with constraints.Ordered then people can't use it with types which don't have < implemented
// so prefer writing a function rather than writing a constraint that requires a method
// eg. for both points 1 and 2

type Tree2[T any] struct {
	cmp  func(T, T) int
	root *node[T]
}

type node[T any] struct {
	left, right *node[T]
	data        T
}

func (bt *Tree2[T]) find(val T) **node[T] {
	pl := &bt.root
	for *pl != nil {
		switch cmp := bt.cmp(val, (*pl).data); {
		case cmp < 0:
			pl = &(*pl).left
		case cmp > 0:
			pl = &(*pl).right
		default:
			return pl
		}
	}
	return pl
}

// 3) When a method looks the same for all types
type SliceFn[T any] struct {
	s   []T
	cmp func(T, T) bool
}

func (s SliceFn[T]) Len() int { return len(s.s) }
func (s SliceFn[T]) Swap(i, j int) {
	s.s[i], s.s[j] = s.s[j], s.s[i]
}

func (s SliceFn[T]) Less(i, j int) bool {
	return s.cmp(s.s[i], s.s[j])
}

func SortFn[T any](s []T, cmp func(T, T) bool) {
	sort.Sort(SliceFn[T]{s, cmp})
}

//* What are type parameters not useful
// 1) When just calling a method on the type argument

// eg.
// good
// func ReadFour(r io.Reader) ([]byte, error)

// bad
// func ReadFour[T io.Reader](r T) ([]byte, error)

// 2) When the implementation of a common method is different for each type
// eg. Read from file and read from a random number generator

// 3) When the operation is different for each type, even without a method

//* When are type parameters useful?
// 1) Avoid boilerplate
// 1.1) Corollary: don't use type parameters prematurely; wait until you are about to write boilerplate code.
