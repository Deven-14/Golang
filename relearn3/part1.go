package main

import (
	"fmt"
	"math/rand"
)

func add(num1, num2 int) int {
	return num1 + num2
}

func swap(num1, num2 int) (int, int) {
	return num2, num1
}

func double(num int) (doubleNum int) {
	doubleNum = num * 2
	return
}

var c, python, java, golang bool

const (
	no1        = true
	no2 uint64 = 1<<64 - 1
)

func part1() {
	var abc bool = true
	var abc2 = false
	x, y := 1, 2
	x, y = swap(x, y)
	fmt.Println(rand.Intn(10), add(5, 10), x, y, double(5), abc, abc2, no1, no2, float32(x))
}
