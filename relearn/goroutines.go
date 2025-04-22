package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for range 5 {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func channels() {
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(s[:len(s)/2], c) // each goroutine doing half the work
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // recieve from c

	fmt.Println(x, y, x+y)
}

func bufferedChannels() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
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
	c := make(chan int, 10)
	go fibonacci2(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}

func fibonacci3(c, quit chan int) {
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

func multiChannelsWaitUsingSelect() {
	// The select statement lets a goroutine wait on multiple communication operations.
	// A select blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready.

	c := make(chan int)
	quit := make(chan int)
	go func() {
		for range 10 {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci3(c, quit)
}

func selectDefault() {

	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!.")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}

}

func goroutinues() {
	go say("hello")
	say("world")

	// channels()

	// bufferedChannels()

	// rangeAndClose()

	// multiChannelsWaitUsingSelect()

	selectDefault()

}
