package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// contraints.Ordered
// type Ordered interface {
// 	Integer|Float|~string
// }
// ~string is the set of types that consists of all types whose underlying type is string

// constraint literals
// [S interface{~[]E}, E interface{}]
// [S ~[]E, E interface{}]
// [S ~[]E, E any]

// making types / alias
// type Point []int32

// This implementation has a problem
func ScaleOld[E constraints.Integer](s []E, c E) []E {
	r := make([]E, len(s))
	for i, v := range s {
		r[i] = v * c
	}
	return r
}

type Point []int32 // multi dimensional point

func (p Point) String() string {
	// details not important
	return "haha"
}

// func ScaleOldAndPrint(p Point) {
// 	r := ScaleOld(p, 2)
// 	fmt.Println(r.String()) // compiler throws error for r.String() coz return type []E i.e., []int32 doesn't have String() method, but Point has String method and it's underlying type is []int32, so do the below changes to make it work
// }

// Compiler error:
// r.String undefined
// (type []int32 has no field or method String)

func Scale[S ~[]E, E constraints.Integer](s S, c E) S {
	r := make(S, len(s))
	for i, v := range s {
		r[i] = v * c
	}
	return r
}

// now this allows the return type to be of any type whose underlying type is []int32

func ScaleOldAndPrint(p Point) {
	r := Scale(p, 2)
	fmt.Println(r.String())
}
