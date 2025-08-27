package main

import (
	"fmt"
	"math"
	"strings"
)

func pointers() {
	var p *int
	x := 5
	p = &x
	fmt.Println(x, *p)
	*p = 12
	fmt.Println(x, *p)
}

type vertex struct {
	x, y int
}

func structs() {

	fmt.Println(vertex{1, 2})
	v := vertex{1, 2}
	fmt.Println(v.x, v.y)

	p := &v
	p.x = 5
	fmt.Println(*p)

	// sturct literals
	v2 := vertex{x: 7}
	fmt.Println(v2)

}

func arrays() {

	var a [10]int
	a = [10]int{1, 2, 3, 4}
	fmt.Println(a)

}

func slices() {

	a := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := a[1:4]
	fmt.Println(a, s)

	// slices change arrays, slices are just references to section of underlying arrays
	// slices does not store any data
	s[0] = 10
	fmt.Println(a, s)

	// slice literals
	s2 := []bool{true, true, false}
	s3 := []struct {
		x, y int
	}{
		{2, 1}, {3, 2}, {4, 5},
	}
	fmt.Println(s2, s3)

	var s4 []float64 // nil is 0 value
	fmt.Println(s4, s4 == nil, len(s4), cap(s4))

	s5 := make([]int, 5)     // len=5, cap=5
	s6 := make([]int, 0, 10) // len=0, cap=10

	s5 = append(s5, 11) // [0, 0, 0, 0, 0, 11], len=6, cap=10 cap doubles
	s6 = append(s6, 11) // [11] as len was 0, len=1, cap=10
	// important to notice that other elements are not displayed as 0, coz array is considered only till its length
	fmt.Println(s5, len(s5), cap(s5), s6, len(s6), cap(s6))

	s5 = s5[:cap(s5)]
	fmt.Println(s5, len(s5), cap(s5))
	s5 = s5[1:]
	fmt.Println(s5, len(s5), cap(s5))

	// slices of slices (slices can contain any type including itself)
	board := [][]string{
		{"ab", "cd", "de"},
		{"ef", "gh", "ij"},
	}
	fmt.Println(board)

	for i, v := range s5 {
		fmt.Print(i, " - ", v, ", ")
	}
	fmt.Println()
	for i := range s5 {
		fmt.Print(i, ", ")
	}
	fmt.Println()

}

func maps() {

	var m map[string]vertex
	m = make(map[string]vertex)
	m["abc"] = vertex{1, 2}
	fmt.Println(m)

	// map literals
	m2 := map[string]vertex{
		"abc": {1, 2},
		"def": {3, 4},
	}
	fmt.Println(m2)

	m2["ghi"] = vertex{5, 6}
	fmt.Println(m2["ghi"])
	delete(m2, "ghi")
	elem, ok := m2["ghi"]
	fmt.Println(elem, ok) // If key is not in the map, then elem is the zero value for the map's element type.

	m3 := wordCount("abc def   gehi adikafhjfsd alsdkfjhas asldkfj abc def")
	fmt.Println(m3)

}

func wordCount(s string) map[string]int {
	m := map[string]int{}
	for _, word := range strings.Fields(s) {
		m[word] += 1
	}
	return m
}

func compute34(fun func(float64, float64) float64) float64 {
	return fun(3, 4)
}

func functionValues() {
	// Functions are values too. They can be passed around just like other values.

	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))
	fmt.Println(compute34(hypot))
	fmt.Println(compute34(math.Pow))

}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func fibonacci() func() int {
	first, second := 0, 1
	return func() int {
		prev := first
		first, second = second, first+second
		return prev
	}
}

func functionClosures() {
	pos, neg := adder(), adder()
	for i := range 10 {
		fmt.Print(pos(i), neg(-2*i), ", ")
	}
	fmt.Println()

	f := fibonacci()
	for range 10 {
		fmt.Println(f())
	}
}

func part3() {

	pointers()

	structs()

	arrays()

	slices()

	maps()

	functionValues()

	functionClosures()

}
