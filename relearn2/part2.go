package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	x, y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

type MyFloat float64

func (f *MyFloat) sqrt() (float64, error) {
	if *f < 0 {
		return float64(0), ErrNegativeSqrt(*f)
	}
	return math.Sqrt(float64(*f)), nil
}

type IMyFloat interface {
	sqrt() (float64, error)
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot sqrt negative numbers: %v", float64(e)) // float64(e) is important to avoid infinite loop
}

func methods() {
	v := Vertex{1.5, 2.6}
	fmt.Println(v.Abs())

	f := MyFloat(5.5)
	fmt.Println(f.sqrt())

	var fi IMyFloat = &f
	fmt.Println(fi.sqrt())

	// type assertions
	var i any = "deven"
	s := i.(string)   // test and then assign to s, if test fails then panic
	s2, ok := i.(int) // no panic because of ok.
	fmt.Println(s, s2, ok)

	// errors
	f2 := MyFloat(-5.5)
	sqrt, err := f2.sqrt()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(sqrt)
	}

}

func part2() {

	methods()

}
