package main

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"time"

	"strings"

	"golang.org/x/tour/pic"
	"golang.org/x/tour/wc"
)

func add(x, y int) int {
	return x + y
}

func swap(x, y int) (int, int) {
	return y, x
}

const a = 5
const (
	Big   = 1 << 100 // cannot print, value too huge
	Small = Big >> 99
)

func sumOfN1(n int) (sum int) {
	for i := 0; i < n; i++ {
		sum += i
	}
	sum = 0
	// new way
	for j := range n {
		sum += j
	}
	return sum
}

func sumOfN2(n int) (sum int) {
	i := 0
	for i < n { // for ; i < n;
		sum += i
		i++
	}
	return sum
}

func Sqrt(x float64) float64 {
	//  Newton's method - https://go.dev/tour/flowcontrol/8
	z := 1.0
	prev := 0.0
	for math.Abs(z-prev) > 0.0001 {
		prev = z
		z -= (z*z - x) / (2 * z)
		// fmt.Println(z, prev)
	}
	fmt.Printf("Sqrt of %f = %f\n", x, z)
	return z
}

func Switch() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}
}

func Switch2() {
	fmt.Print("When's Saturday = ")
	today := time.Now().Weekday()

	switch time.Saturday {
	case today:
		fmt.Println("today.")
	case today + 1:
		fmt.Println("tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("far away.")
	}

}

func SwitchTrue() {
	// switchTrue
	t := time.Now()
	switch { // switch true
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

func pointers() {
	x := 5
	var p *int = &x
	fmt.Println(*p, x)
	*p = 5
	fmt.Println(*p, x)
}

func structs() {
	type Vertex struct {
		X int
		Y int
	}

	v := Vertex{1, 2}

	fmt.Println(v, v.X, v.Y)

	p := &v
	p.X = 5
	fmt.Println(p, v, p.X, v.X)

	type Vertex2 struct {
		X, Y int
	}

	var (
		v1 = Vertex{1, 2}
		v2 = Vertex{X: 1}
		v3 = Vertex{}
		p2 = &Vertex{1, 2}
	)
	fmt.Println(v1, v2, v3, p2)
}

func arrays() {

	var a [10]int
	a[0] = 5
	a[1] = 10
	fmt.Println(a)

	b := [3]string{"a", "b", "c"}
	fmt.Println(b)

}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func Pic(dx, dy int) [][]uint8 {
	a := make([][]uint8, dy)
	for i := range a {
		a[i] = make([]uint8, dx)
	}

	for i := range a {
		for j := range a[i] {
			a[i][j] = uint8(i + j)
		}
	}

	return a
}

func slices() {

	a := [5]int{1, 2, 3, 4, 5}
	s := a[1:4]
	s2 := a[0:2]
	fmt.Println(a, s, s2)
	s[0] = 1
	fmt.Println(a, s, s2) // a, s, s2 all change

	c := []bool{true, false, true} // builds an array and then builds a slice that references it
	fmt.Println(c, a[:5], a[2:], a[:])

	s3 := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s3)

	s4 := []int{2, 3, 5, 7, 11, 13}
	printSlice(s4) // len = 6, cap=6

	// Slice the slice to give it zero length.
	printSlice(s4[:0]) // len=0, cap=6

	// Extend its length.
	printSlice(s4[:4]) // len=4, cap=6

	// Drop its first two values.
	printSlice(s4[2:]) // * len=4, cap=4

	printSlice(s4[2:4]) // len=2, cap=4

	var s5 []int                                 // nil
	fmt.Println(s5, len(s5), cap(s5), s5 == nil) // true

	d := make([]int, 5)    // len = cap = 5
	e := make([]int, 0, 5) // len = 0, cap = 5
	printSlice(d)
	printSlice(e)

	f := [][]string{{"a", "b"}, {"c", "d"}}
	fmt.Println(f)

	g := []int{}

	g = append(g, 0)
	g = append(g, 1, 2)
	printSlice(g)

	h := g[:1]
	h = append(h, 5)
	fmt.Println(h, g) // modifies g

	for i, v := range g {
		fmt.Print(i, v, "  ")
	}

	for i := range g { // * we can drop the value if we need only the index, and we can also use _ to skip like i, _ or _, v
		fmt.Print(i, "  ")
	}

	pic.Show(Pic)

}

func WordCount(s string) map[string]int {
	m := map[string]int{}
	for _, word := range strings.Fields(s) {
		m[word] += 1
	}
	return m
}

func maps() {

	type Vertex struct {
		x, y int
	}

	var m map[string]Vertex // map variable is declared and not initialized

	m = map[string]Vertex{}

	m["abc"] = Vertex{1, 2}
	fmt.Println(m)

	m2 := make(map[string]Vertex) // declared and initialized

	m2["def"] = Vertex{3, 4}
	fmt.Println(m2)

	var m3 = map[string]Vertex{
		"a": Vertex{1, 2},
		"b": Vertex{3, 4},
	}
	fmt.Println(m3)

	var m4 = map[string]Vertex{
		"a": {1, 2},
		"b": {3, 4},
	}
	fmt.Println(m4)

	m4["c"] = Vertex{5, 6}
	elem, ok := m4["a"]
	delete(m4, "a")
	fmt.Println(m4, elem, ok)

	wc.Test(WordCount)

}

func compute(fn func(x, y float64) float64) float64 {
	return fn(3, 4)
}

func adder() func(x int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	first, second := 0, 1
	return func() int {
		prev := first
		first, second = second, first+second
		return prev
	}
}

func functionValues() {

	sumOfSquares := func(x, y float64) float64 {
		return math.Pow(x, 2) + math.Pow(y, 2)
	}

	fmt.Println(sumOfSquares(5, 4))

	fmt.Println(compute(sumOfSquares))
	fmt.Println(compute(math.Pow))

	pos, neg := adder(), adder()
	for i := range 10 {
		fmt.Println(pos(i), neg(-2*i))
	}

	f := fibonacci()
	for range 10 {
		fmt.Println(f())
	}

}

func main() {
	var a, b = 1, true
	var x, y = swap(3, 4)
	c := float64(a)

	defer fmt.Println("completed with main function")

	fmt.Println("abc", rand.Intn(10), add(3, 4), a, b, x, y, c, float64(Big), Small)

	fmt.Println(sumOfN1(5), sumOfN2(5))

	if x == 4 {
		fmt.Println(a, b)
	}

	if z := 5; z > 0 {
		fmt.Println(z)
	} else {
		fmt.Println(-z)
	}

	// Sqrt(49)

	// Switch()
	// Switch2()
	// SwitchTrue()

	// pointers()

	// structs()

	// arrays()

	// slices()

	// maps()

	functionValues()

}
