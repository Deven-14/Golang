package main

import (
	"fmt"
	"math"
	"time"
)

type ErrNegativeSqrt float64

// Note: A call to fmt.Sprint(e) inside the Error method will send the program into an infinite loop. You can avoid this by converting e first: fmt.Sprint(float64(e)). Why?
func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprint("cannot Sqrt negative number: ", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z, prev := 1.0, 0.0
	for math.Abs(z-prev) > 0.0001 {
		prev = z
		z -= (z*z - x) / (2 * z)
	}
	return z, nil
}

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func errorsExample() {

	if err := run(); err != nil {
		fmt.Println(err)
	}

	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
