package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func example1() {

	fmt.Println("\nExample1:")
	ch := make(chan int)
	wg.Add(2)
	// based on arrow head we can say where the data is flowing
	go func() {
		i := <-ch // flow data out of ch and into var i
		fmt.Println(i)
		wg.Done()
	}()
	go func() {
		ch <- 42 // flow data into ch
		wg.Done()
	}()
	wg.Wait()

}

func example2() {

	fmt.Println("\nExample2:")
	ch := make(chan int)
	wg.Add(2)
	go func() {
		i := <-ch
		fmt.Println(i)
		wg.Done()
	}()
	go func() {
		i := 42
		ch <- i // copy of data is passed (value type)
		i = 27
		wg.Done()
	}()
	wg.Wait()

}

func example3() {

	fmt.Println("\nExample3:")
	ch := make(chan int)

	for j := 0; j < 5; j++ {
		wg.Add(2)

		go func() {
			i := <-ch
			fmt.Println(i)
			wg.Done()
		}()
		go func() {
			ch <- 42
			wg.Done()
		}()

	}

	wg.Wait()

}

func deadlockExample() {

	fmt.Println("\nDeadlock Example:")
	ch := make(chan int)

	go func() {
		i := <-ch
		fmt.Println(i)
		wg.Done()
	}()

	for j := 0; j < 5; j++ {
		wg.Add(2)

		go func() {
			ch <- 42
			wg.Done()
		}()

	}

	wg.Wait()

	// deadlock because, first one thread is waiting for value and then enters for loop and 5 thread try writes to channel, but only one thread is able to then it exits and the first thread receives the data from the channel and prints it and from the remaining 4 threads one of them writes and remaining 3 are waiting at ch <- 42 to write, but no thread is there to receieve (BY DEFAULT WE ARE WORKING WITH UNBUFFERED CHANNELS AND ONLY ONE MESSAGE CAN BE IN THE CHANNEL AT ONE TIME) and the runtime detects there's a deadlock since there's no one to receive and throws error

}

func example4() {

	fmt.Println("\nExample4:")
	ch := make(chan int)

	wg.Add(2)

	// both functions are acting both as readers and writers but usually we going to need dedicated functions for reading and writing, soln in eg5

	go func() {
		i := <-ch
		fmt.Println(i)
		ch <- 27
		wg.Done()
	}()
	go func() {
		ch <- 42
		fmt.Println(<-ch)
		wg.Done()
	}()

	wg.Wait()

}

func example5() {

	fmt.Println("\nExample5, Resticting Data Flow:")
	ch := make(chan int)

	wg.Add(2)

	// send only channel
	go func(ch <-chan int) {
		i := <-ch
		fmt.Println(i)
		// ch <- 27 // shows complier error coz data type is <-chan, making this function a send only
		wg.Done()
	}(ch)

	// receive only channel
	go func(ch chan<- int) {
		ch <- 42
		// fmt.Println(<-ch) // shows complier error coz data type is chan<-, making this function a read only
		wg.Done()
	}(ch)

	wg.Wait()

}

func example6() {

	fmt.Println("\nExample6:")
	ch := make(chan int)

	wg.Add(2)

	go func(ch <-chan int) {
		i := <-ch
		fmt.Println(i)
		wg.Done()
	}(ch)

	go func(ch chan<- int) {
		ch <- 42
		ch <- 27 // deadlock, (BY DEFAULT WE ARE WORKING WITH UNBUFFERED CHANNELS AND ONLY ONE MESSAGE CAN BE IN THE CHANNEL AT ONE TIME), so one message goes into the channel and that is read and another message goes into the channel and no one to read it, so deadlock, soln ep7, buffered channels
		wg.Done()
	}(ch)

	wg.Wait()

}

func example7() {

	fmt.Println("\nExample7:")
	ch := make(chan int, 50) // buffered channel, tells channel to create an internal datastore that can store 50 integers

	// buffered channels help solve the issue of unbuffered channels i.e., 1) Channels block sender side till receiver is available 2) block receiver side till message is available

	// * buffered channels are useful when sender and receiver have assymmetric loading, i.e., if we send messages faster than we can receive them

	wg.Add(2)

	go func(ch <-chan int) {
		i := <-ch
		fmt.Println(i)
		wg.Done()
	}(ch)

	go func(ch chan<- int) {
		ch <- 42
		ch <- 27 // the channel accepts this message but it gets lost as no one is there to read it
		wg.Done()
	}(ch)

	// 42 and 27 are put into channel but since there is only one read from channel, only one message 42 is read and the other is lost

	wg.Wait()

}

func example8() {

	fmt.Println("\nExample8:")
	ch := make(chan int, 50)

	wg.Add(2)

	go func(ch <-chan int) {
		i := <-ch
		fmt.Println(i)
		i = <-ch
		fmt.Println(i)
		wg.Done()
	}(ch)

	go func(ch chan<- int) {
		ch <- 42
		ch <- 27
		wg.Done()
	}(ch)

	// 2 msgs are put into the channel and both of them are read, no data lost.
	// now lets say we have 50+ (n) burst messages being sent and we don't know how many message to receive then we can use the 'range' syntax in eg9
	wg.Wait()

}

func example9() {

	fmt.Println("\nExample9:")
	ch := make(chan int, 50)

	wg.Add(2)

	go func(ch <-chan int) {
		for i := range ch {
			fmt.Println(i)
		}
		wg.Done()
	}(ch)

	go func(ch chan<- int) {
		ch <- 42
		ch <- 27
		wg.Done()
	}(ch)

	// here receiver is getting deadlocked as only 2 messages in sent to the buffer but we are continuing to montinor for additional messages after reading those 2 messages, so receiver thread is deadlocked

	wg.Wait()

}

func example10() {

	fmt.Println("\nExample10:")
	ch := make(chan int, 50)

	wg.Add(2)

	go func(ch <-chan int) {
		for i := range ch {
			fmt.Println(i)
		}
		wg.Done()
	}(ch)

	go func(ch chan<- int) {
		ch <- 42
		ch <- 27
		close(ch)
		wg.Done()
	}(ch)

	// there can be an infinite messages in channel as we can keep sending messages into the channel, so to stop the for 'range' loop from monitoring for more messages, we can close the channel and the for 'range' loop detects that and stops monitoring for more messages

	// ONCE THE CHANNEL IS CLOSED IT CANNOT BE OPENED AGAIN

	wg.Wait()

}

func example11() {

	fmt.Println("\nExample11:")
	ch := make(chan int, 50)

	wg.Add(2)

	go func(ch <-chan int) {

		// same as for 'range' loop, but the ', ok' syntax is used best used when we are not using a loop and just trying to receive messages and we don't know if message is sent and then we can use ', ok' syntax to check if there are any messages then recieve it...

		// and as u can see with ', ok' syntax and the for loop it goes into infinte loop as wg.Done() is not reachable code

		// for {
		// 	if i, ok := <-ch; ok {
		// 		fmt.Println(i)
		// 	}
		// }
		// wg.Done()

		// so instead this syntax should be used in this manner coz we are having a different function to be executed for each iteration (or if same function then we can use for 'range') and note that since we are using using the ', ok' syntax, the last receiver is notified that the channel is close and it didn't execute and it didn't wait for a message and didn't deadlock

		if i, ok := <-ch; ok {
			fmt.Println(i, 1)
		}

		if i, ok := <-ch; ok {
			fmt.Println(i, 2)
		}

		if i, ok := <-ch; ok {
			fmt.Println(i, 3)
		}

		wg.Done()

	}(ch)

	go func(ch chan<- int) {
		ch <- 42
		ch <- 27
		close(ch)
		wg.Done()
	}(ch)

	// there can be an infinite messages in channel as we can keep sending messages into the channel, so to stop the for 'range' loop from monitoring for more messages, we can close the channel and the for 'range' loop detects that and stops monitoring for more messages

	// ONCE THE CHANNEL IS CLOSED IT CANNOT BE OPENED AGAIN

	wg.Wait()

}

func example12() {

	fmt.Println("\nExample12:")
	ch := make(chan int)

	wg.Add(2)

	go func(ch <-chan int) {
		for i := range ch {
			fmt.Println(i)
		}
		wg.Done()
	}(ch)

	go func(ch chan<- int) {
		ch <- 42
		ch <- 27
		close(ch)
		wg.Done()
	}(ch)

	// In this case the sender is blocked until the receiver processes it and after that the sender can send the next message, so the sender is blocked until the receiver is ready to receive the message

	wg.Wait()

}

func main() {

	// * channels are mainly used in the context of goroutines coz channels are designed to synchronize data transmission between multiple go routines

	// to create a channel only make() can be used
	// the type of channel is very strongly typed, int then we can pass only int, if pointer to integers then we can send only pointer to integers,
	// channel can only receieve and send messages of that type

	example1()

	example2()

	// use case of goroutines
	// if u have data that is asynchronously processed, so you can generate the data very, very quickly, but it takes time to process it. Or maybe it takes long time to generate that data so you've got multiple generators but it can processed very quickly. So we might want to have a different number of goroutines that are sending data into a channel then you have receiving.

	example3()

	// deadlockExample()

	example4()

	example5()

	// buffered channels

	// example6() // deadlock example

	example7()

	example8()

	// example9() // receiver deadlock

	example10()

	example11()

	example12()

}
