package main

import (
	"fmt"
	"math/rand"
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

func main() {
	var a, b = 1, true
	var x, y = swap(3, 4)
	c := float64(a)
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

}
