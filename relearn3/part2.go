package main

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"time"
)

func loops() {

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	for i := range 10 {
		fmt.Println(i)
	}

	sum := 1
	for sum < 100 {
		sum += sum
	}
	fmt.Println(sum)

	sum = 1
	for {
		if sum >= 100 {
			break
		}
		sum += sum
	}
	fmt.Println(sum)

}

func ifConditions() {

	if x := rand.Intn(10); x >= 5 {
		fmt.Println(math.Sqrt(float64(x)))
	} else {
		fmt.Println(x)
	}

}

func switchStatement() {

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("mac")
	case "linux":
		fmt.Println("linux")
	default:
		fmt.Println("other", os)
	}

	t := time.Now()
	switch { // switch true, checks for the case which is true
	case t.Hour() < 12:
		fmt.Println("Good Morning")
	case t.Hour() < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening")
	}

	switch false {
	case t.Hour() < 12:
		fmt.Println("Not Good Morning")
	case t.Hour() < 17:
		fmt.Println("Not Good afternoon")
	default:
		fmt.Println("Not Good evening")
	}

}

func part2() {
	defer fmt.Println("world")
	fmt.Println("hello")

	loops()

	ifConditions()

	x := 2.0
	Sqrt1(x)
	Sqrt2(x)

	switchStatement()

}
