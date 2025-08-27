package main

import (
	"fmt"
	"math"
)

func Sqrt1(x float64) float64 {
	z := 1.0
	prev := 0.0
	for math.Abs(z-prev) > 0.0001 {
		prev = z
		z -= (z*z - x) / (2 * z)
		fmt.Println(z, prev)
	}
	fmt.Printf("Sqrt of %f = %f\n", x, z)
	return z
}

func Sqrt2(x float64) float64 {

	z := 1.0
	for prev := 0.0; math.Abs(z-prev) > 0.0001; z -= (z*z - x) / (2 * z) {
		prev = z
	}

	fmt.Printf("Sqrt of %f = %f\n", x, z)
	return z
}
