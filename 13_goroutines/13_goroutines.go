package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func sayHello() {
	fmt.Println("Hello")
}

func example1() {
	fmt.Println("\nExample1:")
	go sayHello()
	time.Sleep(100 * time.Microsecond) // bad practice to wait for it, soln in eg5
}

func example2() {
	fmt.Println("\nExample2:")
	var msg = "Hello"
	go func() {
		fmt.Println(msg) // goroutine takes care of using the msg variable from outside (created a dependancy between a variable in the goroutine and a variable in the main function), but this can create a problem see eg3
	}()
	time.Sleep(100 * time.Millisecond)
}

func example3() {
	fmt.Println("\nExample3:")
	var msg = "Hello"
	go func() {
		fmt.Println(msg)
	}()
	msg = "Goodbye" // goodbye is printed, coz of race condition
	time.Sleep(100 * time.Millisecond)
	// fix in eg4
}

func example4() {
	fmt.Println("\nExample4:")
	var msg = "Hello"
	go func(msg string) {
		fmt.Println(msg)
	}(msg)
	msg = "Goodbye"
	time.Sleep(100 * time.Millisecond)
}

// For syncronization
var wg = sync.WaitGroup{}

func example5() {

	fmt.Println("\nExample5:")
	var msg = "Hello"

	wg.Add(1) // saying 1 goroutine (added to counter)
	go func(msg string) {
		fmt.Println(msg)
		wg.Done() // decrements the counter by 1
	}(msg)

	msg = "Goodbye"
	wg.Wait() // waits for counter to become 0

}

// Race Condition Example
var counter = 0

func sayHello2() {
	fmt.Printf("Hello #%v\n", counter)
	wg.Done()
}

func increment() {
	counter++
	wg.Done()
}

func RaceConditionExample() {
	fmt.Println("\nRace Condition Example:")
	for i := 0; i < 10; i++ {
		wg.Add(2)
		go sayHello2()
		go increment()
	}
	wg.Wait()
	// order is wrong, some values are double or more and some are missing
}

// Mutex Example
var m = sync.RWMutex{}

func sayHello3() {
	m.RLock()
	fmt.Printf("Hello #%v\n", counter)
	m.RUnlock()
	wg.Done()
}

func increment3() {
	m.Lock()
	counter++
	m.Unlock()
	wg.Done()
}

func PartialRCFixExample() {
	fmt.Println("\nPartial Fix For Race Condition Example:")
	for i := 0; i < 10; i++ {
		wg.Add(2)
		go sayHello3()
		go increment3()
	}
	wg.Wait()
	//problem here is 2 or more sayHello might get executed before increment function can write, so here the order is correct but same values are printed many times and some are missing
}

// Complete fix
func sayHello4() {
	fmt.Printf("Hello #%v\n", counter)
	m.RUnlock()
	wg.Done()
}

func increment4() {
	counter++
	m.Unlock()
	wg.Done()
}

func CompleteRCFixExample() {
	fmt.Println("\nComplete Fix For Race Condition Example:")
	for i := 0; i < 10; i++ {
		wg.Add(2)
		m.RLock()
		go sayHello4()
		m.Lock()
		go increment4()
	}
	wg.Wait()
}

func MutexExample() {

	// A Simple Mutex is simply locked or unlocked and so if the Mutex is locked and something tries to manipulate that value, it has to wait until the mutex is unlocked and then can obtain the mutex lock itself

	// so we can protect parts of our code so that only one entity can be manipulating that code at a time.

	// A RW Mutex - as many things can want to can read this but only one can write it at a time and if any one if reading then we can't write to it at all.

	counter = 0
	PartialRCFixExample()

	counter = 0
	CompleteRCFixExample()

	// but after complete fix, this runs as a single thread operation and the whole point of goroutines is actually gone

}

func main() {
	// using green threads, these don't use massive heavy threads but an abstraction of a thread called goroutine

	// inside go runtime there is a scheduler that's going to map these goroutines onto these operating system threads for periods of time and the scheduler will take turns with every CPU thread available and assign the different goroutines a certain amount of processing time on those threads (we don't have to interact with those low level threads we're interacting with these high level goroutines) and since we have the abstration goroutines can start with very very small stack spaces because they can be reallocated very very quickly and so they're very cheap to create and to destroy.

	// very common to have tens of thousands of goroutines running at the same time and the application has no problem with that

	// other languages which rely operating system threads that have huge (1MB ram) of overhead, there's no way of running 10000 threads in an environment

	example1()

	example2()

	example3()

	example4()

	// waitgroup - for synchronization
	example5()

	RaceConditionExample()

	MutexExample()

	// * For Parallelism
	fmt.Printf("Threads: %v\n", runtime.GOMAXPROCS(-1))

	// running applications in a single threaded way
	runtime.GOMAXPROCS(1) // useful in situations where there is a lot of data synchronizations going on and to avoid any kind of race conditions that parallelism can incur (but that's kinda a architecture problem)

	fmt.Printf("Threads: %v\n", runtime.GOMAXPROCS(-1))

	runtime.GOMAXPROCS(100)

	fmt.Printf("Threads: %v\n", runtime.GOMAXPROCS(-1))

	// general advice - atleast one operating system thread per core is a minimum but a lot of times it is seen that the application will get faster by increasing go max procs beyond that value.
	// but if you get up too high, eg. 100, then you can run into other problems coz now you've got additional memory overhead coz you're maintaining 100 operating system threads and you're scheduler has to work harder, so eventually the performance peaks and then it begins to fall back off coz the application is constantly rescheduling goroutines on different threads and losing time every time that occurs.

	// in production - before pushing it to production, perform a performance test suite with varying values of GOMAXPROCS to see where it's going to perform the best

	// best practices

	// 1
	// Don't create goroutines in libraries
	// let consumers control concurrency
	// but in case if the func called will return a channel which will return a result then having the goroutine in there might not be such a bad thing coz ur consumer never really has to worry about how that unit of work is getting done

	// 2
	// when creating a goroutine, know how it will end
	// avoid subtle memeory leaks

	// 3
	// check for race condition at compile time
	// go run -race main.go

}
