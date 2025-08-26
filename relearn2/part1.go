package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func add(x, y int) (sum int) {
	sum = x + y
	return
}

func swap(x, y string) (string, string) {
	return y, x
}

var ABC = "deven"

// function values
func compute(fn func(int, int) int) int {
	return fn(3, 4)
}

// function closures
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func basics() {
	a, b := swap("a", "b")
	fmt.Println("something", rand.Intn(25), add(10, 15), a, b)

	var name string
	name = "deven"
	var name2 = "deven"
	name3 := "deven"
	var num1, num2 = 1, 2
	const World = "abc"
	fmt.Println(name, name2, name3, num1, num2, float32(num1), World)

	for i := range 10 {
		fmt.Print(i)
	}

	if num3 := 5; num3 > 0 {
		defer fmt.Println("num3")
	}

	today := time.Now().Weekday()
	switch time.Saturday {
	case today:
		fmt.Println("today")
	case today + 1:
		fmt.Println("tomorrow")
	default:
		fmt.Println("far")
	}

	var v int = 5
	var p *int = &v
	fmt.Println(v, *p)

	type Vertex struct {
		x, y int
	}

	v1 := Vertex{1, 2}
	p1 := &v1
	fmt.Println(Vertex{1, 2}, v1, p1.x, Vertex{y: 1}, Vertex{})

	var a2 [10]int = [10]int{1, 2, 3, 4}
	d := []int{1, 2, 3, 4}
	e := []struct {
		i bool
		s string
	}{
		{true, "a"},
		{false, "b"},
	}
	f := make([]int, 0, 10)
	f = append(f, 5, 10, 15)
	fmt.Println(a2, a2[3:6], d, e, f, f[1:], f[1:cap(f)], [][]string{[]string{"1", "2"}})

	for i, v := range f {
		fmt.Println(i, v)
	}

	m := map[int]string{}
	m2 := make(map[int]string)
	m[1] = "a"
	m2[2] = "b"
	fmt.Println(m, m2)
	delete(m2, 2)
	ele, ok := m[1]

	fmt.Println(ele, ok)

	m3 := map[string]int{}
	s := "abadc daflkdj alsdifh kalhf2893yrhn a9op;8rfyu30sioxfh aosid   "
	for _, word := range strings.Fields(s) {
		m3[word] += 1
	}
	fmt.Println(m3)

	sum := func(x, y int) int {
		return x + y
	}
	fmt.Println(compute(sum))

	pos, neg := adder(), adder()
	for i := range 5 {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)

	}

}

func part1() {
	basics()
}
