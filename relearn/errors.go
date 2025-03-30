package main

import (
	"fmt"
	"math"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

// error interface in built-in
// type error interface {
//     Error() string
// }

func runError() error {
	return &MyError{time.Now(), "It didn't work"}
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt2(x float64) (float64, error) {
	if x < 0 {
		return float64(0), ErrNegativeSqrt(x)
	}
	z := 1.0
	prev := 0.0
	for math.Abs(z-prev) > 0.0001 {
		prev = z
		z -= (z*z - x) / (2 * z)
		// fmt.Println(z, prev)
	}
	return z, nil
}

func runErrors2() {
	fmt.Println(Sqrt2(2))
	fmt.Println(Sqrt2(-2))
}
