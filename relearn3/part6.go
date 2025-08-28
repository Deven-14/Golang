package main

import (
	"fmt"
	"time"
)

// * A goroutine is a lightweight thread managed by the Go runtime.

func say(s string) {

	for range 5 {
		time.Sleep(5 * time.Millisecond)
		fmt.Println(s)
	}
}

func sum(s []int, c chan int) {
	total := 0
	for _, v := range s {
		total += v
	}
	c <- total
}

func channels() {

	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)

}

func bufferedChannels() {

	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch, <-ch)

}

func fibonacci2(n int, c chan int) {
	x, y := 0, 1
	for range n {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func rangeAndClose() {

	ch := make(chan int, 10)
	go fibonacci2(30, ch)
	for i := range ch {
		fmt.Print(i, ", ")
	}
	fmt.Println()

	// ! v, ok := <-ch

}

func fibonacci3(c chan int, quit chan struct{}) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func selectStatement() {

	// The select statement lets a goroutine wait on multiple communication operations.
	c := make(chan int)
	quit := make(chan struct{})
	go func() {
		for range 10 {
			fmt.Println(<-c)
		}
		quit <- struct{}{}
	}()
	fibonacci3(c, quit)

	// example
	start := time.Now()
	elapsed := func() time.Duration {
		return time.Since(start).Round(time.Millisecond)
	}
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Printf("[%6s] tick.\n", elapsed())
		case <-boom:
			fmt.Printf("[%6s] boom.\n", elapsed())
			return
		default:
			fmt.Println("....", elapsed())
			time.Sleep(50 * time.Millisecond)
		}
	}

}

func part6() {

	go say("world")
	say("hello")

	channels()

	bufferedChannels()

	rangeAndClose()

	selectStatement()

	IsBSTSame()

	mutexExample()

}
