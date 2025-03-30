package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	x, y float64
}

func (v Vertex) abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func (v *Vertex) Scale(f float64) {
	v.x *= f
	v.y *= f
}

func abs(v Vertex) float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func Scale(v *Vertex, f float64) {
	v.x *= f
	v.y *= f
}

type MyFloat float64

func (f MyFloat) abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func methods() {
	v := Vertex{3, 4}
	fmt.Println(v.abs())
	fmt.Println(abs(v))

	f := MyFloat(5.5)
	fmt.Println(f.abs())

	v.Scale(5.5)
	fmt.Println(v)
	Scale(&v, 5.5)
	fmt.Println(v)

	p := &v
	p.Scale(5.5) // p.Scale or v.Scale both work
	fmt.Println(p, v)
	// Scale(v, 5.5) won't work, has to be Scale(&v, 5.5) or Scale(p, 5.5)

	fmt.Println(p.abs()) // p.abs() or v.abs() both will work
	// abs(p / &v) won't work, has to be abs(v) or abs(*p)

}

func relearn2() {

	methods()

}
