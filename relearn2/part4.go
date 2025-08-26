package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for range 5 {
		fmt.Println(s)
		time.Sleep(100 * time.Millisecond)
	}
}

func example1() {
	go say("hello")
	say("world")
}

// channels
// By default, sends and receives block until the other side is ready. This allows goroutines to synchronize without explicit locks or condition variables.
// The example code sums the numbers in a slice, distributing the work between two goroutines. Once both goroutines have completed their computation, it calculates the final result.

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for range n {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}

func fibonacci2(ch, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case ch <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func channels_example() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c

	fmt.Println(x, y, x+y)

	// buffered channels
	// Sends to a buffered channel block only when the buffer is full. Receives block when the buffer is empty.
	ch := make(chan int, 100)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch, <-ch)

	// range and close
	// A sender can close a channel to indicate that no more values will be sent. Receivers can test whether a channel has been closed by assigning a second parameter to the receive expression: after
	// v, ok := <-ch
	// The loop for i := range c receives values from the channel repeatedly until it is closed.
	// Note: Only the sender should close a channel, never the receiver. Sending on a closed channel will cause a panic.
	// Another note: Channels aren't like files; you don't usually need to close them. Closing is only necessary when the receiver must be told there are no more values coming, such as to terminate a range loop.
	ch2 := make(chan int, 10)
	go fibonacci(cap(ch2), ch2) // * offloading the heavy task to some other thread
	for i := range ch2 {
		fmt.Println(i)
	}

	// The select statement lets a goroutine wait on multiple communication operations.
	// * taking the input from the main goroutine and processing in another goroutine
	ch3 := make(chan int)
	quit := make(chan int)
	go func() {
		for range 10 {
			fmt.Println(<-ch3)
		}
		quit <- 0
	}()
	fibonacci2(ch3, quit)

	// The default case in a select is run if no other case is ready.
	// Use a default case to try a send or receive without blocking:
	start := time.Now()
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	elapsed := func() time.Duration {
		return time.Since(start).Round(time.Millisecond)
	}

	for {
		select {
		case <-tick:
			fmt.Println("tick ", elapsed())
		case <-boom:
			fmt.Println("boom ", elapsed())
			return
		default:
			fmt.Println("waiting...", elapsed())
			time.Sleep(50 * time.Millisecond)
		}
	}

}

func part4() {
	example1()
	channels_example()

}
