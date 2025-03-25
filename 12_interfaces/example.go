package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

func example1() {

	// * example why all methods on a given type should have either value or pointer receivers, but not a mixture of both. (We'll see why over the next few pages.)

	// A value of interface type can hold any value that implements those methods.

	// Note: There is an error in the example code on line 29. Vertex (the value type) doesn't implement Abser because the Abs method is defined only on *Vertex (the pointer type).

	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // a MyFloat implements Abser
	a = &v // a *Vertex implements Abser

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	// a = v

	fmt.Println(a.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
