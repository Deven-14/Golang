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

func (v *Vertex) action() float64 {
	// something
	return 5.5
}

// Stringer interface
func (v Vertex) String() string {
	return fmt.Sprintf("Vertex{x: %v, y: %v}", v.x, v.y)
}

// methods are functions
func abs(v Vertex) float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func Scale(v *Vertex, f float64) {
	v.x *= f
	v.y *= f
}

// methods can be for types which are not structs
type MyFloat float64

func (f MyFloat) abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (f MyFloat) action() float64 {
	// something
	return 5.5
}

// interfaces
type Actioner interface {
	action() float64
}

func printAny(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
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

func interfaces() {
	v := Vertex{3, 4}
	f := MyFloat(5.5)

	var a Actioner
	// a = v  // Vertex doesn't implements Actioner, shows error
	a = &v // *Vertex implements Actioner
	fmt.Printf("(%v, %T)\n", a, a)

	a = f // MyFloat implements Actioner
	fmt.Printf("(%v, %T)\n", a, a)
	a = &f // *MyFloat implements Actioner (MyFloat implements, works for *MyFloat implicitely, but no vice versa)
	fmt.Printf("(%v, %T)\n", a, a)

	a.action()

	var f2 MyFloat // only declared and not initialized
	a = f2
	fmt.Printf("(%v, %T)\n", a, a)
	a.action()

	// var a2 Actioner
	// a2.action() // throws runtime error

	printAny(1)
	printAny(5.5)
	printAny("string")

}

func typeAssertions(i interface{}) {
	s := i.(MyFloat) // if i is not MyFloat, then program panics
	fmt.Println(s)

	t, ok := i.(string) // no panic because of ok
	fmt.Println(t, ok)

	u, ok := i.(float64)
	fmt.Println(u, ok)
}

func typeSwitches(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("int", v)
	case string:
		fmt.Println("string", v)
	default:
		fmt.Println("no clue")
	}
}

type IPAddr [4]byte

func (ip IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
}

func relearn2() {

	methods()

	interfaces()

	typeAssertions(MyFloat(5.5))

	typeSwitches("hello")

	// stringer
	ip := IPAddr{1, 2, 3, 4}
	fmt.Println(ip)

}
