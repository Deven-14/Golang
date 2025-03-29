package main

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"time"
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

	Sqrt(49)

	Switch()
	Switch2()
	SwitchTrue()

	pointers()

	structs()

}
