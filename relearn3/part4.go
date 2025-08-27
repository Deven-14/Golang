package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"strings"
)

type Vertex struct {
	x, y float64
}

func (v Vertex) abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func (v *Vertex) scale(f float64) {
	if v == nil {
		fmt.Println("nil")
		return
	}
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

	v := Vertex{1, 2}
	fmt.Println(v.abs())

	f := MyFloat(-5.5)
	fmt.Println(f.abs())

	v.scale(5)
	fmt.Println(v)

	p := &v
	fmt.Println(p, p.abs())
	p.scale(5)
	fmt.Println(p)

	// In general, all methods on a given type should have either value or pointer receivers, but not a mixture of both.

}

type Abser interface {
	abs() float64
}

type Scaler interface {
	scale(float64)
}

func interfaces() {
	// An interface type is defined as a set of method signatures.

	// A value of interface type can hold any value that implements those methods.

	var a Abser
	v := Vertex{-1, 2}
	f := MyFloat(-5)

	a = f // works
	fmt.Println(a.abs())
	fmt.Printf("(%v, %T)\n", a, a)

	a = v // works
	fmt.Println(a.abs())
	fmt.Printf("(%v, %T)\n", a, a)

	a = &v
	// fmt.Println(a.abs()) // doesn't work because *Vertex doesn't implement abs(), Vertex implements abs()
	// because of this, all the methods should be either in value or pointer receivers and not mix
	fmt.Printf("(%v, %T)\n", a, a)

	var v2 *Vertex // not initialized
	var s Scaler = v2
	s.scale(5) // prints null

	var i interface{}
	var i2 any
	i = 5
	i2 = 10
	fmt.Println(i, i2)

}

func typeAssertions() {

	var i any = "string"
	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	// f := i.(float64) // panic

}

func typeSwitches(i any) {

	switch v := i.(type) {
	case int:
		fmt.Println(v, v*2)
	case string:
		fmt.Println(v, v+v)
	default:
		fmt.Println("I don't know the type")
	}

}

type Person struct {
	name string
	age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.name, p.age)
}

func Stringers() {

	p := Person{"Deven", 23}
	fmt.Println(p)

}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot sqrt negative numbers: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return float64(0), ErrNegativeSqrt(x)
	}
	return Sqrt2(x), nil
}

func errors() {

	x := -5.0
	y, err := Sqrt(x)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(y)

}

type rot13Reader struct {
	r io.Reader
}

func (rot rot13Reader) Read(b []byte) (int, error) {
	n, err := rot.r.Read(b)
	for i, char := range b {
		if char >= 'A' && char <= 'Z' {
			b[i] = 'A' + (char-'A'+13)%26
		} else if char >= 'a' && char <= 'z' {
			b[i] = 'a' + (char-'a'+13)%26
		}
	}
	return n, err
}

func readers() {

	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)

}

func part4() {

	methods()

	interfaces()

	typeAssertions()

	typeSwitches(5)
	typeSwitches("6alkdhf")
	typeSwitches(5.33)

	Stringers()

	errors()

	readers()

}
