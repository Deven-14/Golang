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

func main() {
	var a, b = 1, true
	var x, y = swap(3, 4)
	c := float64(a)
	fmt.Println("abc", rand.Intn(10), add(3, 4), a, b, x, y, c, float64(Big), Small)
}
