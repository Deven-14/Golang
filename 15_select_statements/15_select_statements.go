package main

import (
	"fmt"
	"time"
)

const (
	logInfo   = "INFO"
	logWarnig = "WARNING"
	logError  = "ERROR"
)

type logEntry struct {
	time     time.Time
	severity string
	message  string
}

var logCh = make(chan logEntry, 50)

func logger() {

	for entry := range logCh {
		fmt.Printf("%v - [%v]%v\n", entry.time.Format("2006-01-02T15:04:05"), entry.severity, entry.message)
	}

}

// func main() {

// 	// example 1

// * select statements allows goroutine to monitor several channels at once
// 1) block if all channels block
// 2) if muliple channels receive value simultaneously, behaviour is undefined

// 	go logger()

// 	logCh <- logEntry{time.Now(), logInfo, "App is starting"}

// 	logCh <- logEntry{time.Now(), logInfo, "App is shutting down"}

// 	time.Sleep(100 * time.Millisecond)

// 	// here the logCh channel is not closed but still the application shutsdown without deadlock, coz an application shutsdown as soon as the last statement of the main function finishes execution, so everything is torn down and all resources are reclaimed by the OS. What that mean is that our go routine is tron down forcibly and there is no graceful shutdown for this go routine.

// 	// we should have more control over our go routine and WE SHOULD ALWAYS HAVE A STRATEGY FOR HOW YOUR GOROUTINE IS GOING TO SHUTDOWN WHEN YOU CREATE YOUR GOROUTINE. Otherwise, it can be a subtle resource leak, and evetually, it can leak enough resources to bring your application down. examples shown in other files

// }

// func main() {

// 	// example 2

// 	go logger()

// 	defer func() {
// 		close(logCh)
// 	}() // this is perfectly acceptable in this use case, but there is one more method shown in example 3

// 	logCh <- logEntry{time.Now(), logInfo, "App is starting"}

// 	logCh <- logEntry{time.Now(), logInfo, "App is shutting down"}

// 	time.Sleep(100 * time.Millisecond)

// }

var doneCh = make(chan struct{})

// struct with no fields is unique and that it requires zero memeory allocations.
// The intention of this channel is, it can't send any data through except for the fact that a message was sent or received. So is called a Signal only channel.
// There's zero memory allocations and required in sending the message but we do have the ability the receiving side know that a message was sent.
// U might be temped to do var doneCh = make(chan bool) but that actually requires a variable to be 'allocated and copied'. So it is actually better to use an empty struct because it saves a couple of memory allocations. It's a little bit minor but it is something that if you are going to use a channel as a pure message, then you might as well go with the conventions.

func logger2() {

	for {

		select {
		case entry := <-logCh:
			fmt.Printf("%v - [%v]%v\n", entry.time.Format("2006-01-02T15:04:05"), entry.severity, entry.message)
		case <-doneCh:
			fmt.Println("done")
			break
		}

	}

	// this select statement is going to block UNTIL a message is received on one of the channels that is listening for (logCh and doneCh here).

	// if we add a default case then the select statement no longer becomes a blocking select statement, so if there's a message ready on one of the channels that are being monitored, then it's going to execute that code path. If not, it will execute a default block. This is useful if you want to have a  nonblocking select statement.

	// * the doneCh can be passed as a parameter instead of a global variable. Not just the doneCh channel, where ever it is possible use paramters instead of global variables
	// * channels are reference type

}

func main() {

	// example 3

	go logger2()

	logCh <- logEntry{time.Now(), logInfo, "App is starting"}

	logCh <- logEntry{time.Now(), logInfo, "App is shutting down"}

	doneCh <- struct{}{} // type struct{} - struct without fields and the next {} is for initializing the struct

	// you will send the doneCh when you want to kill the goroutine

	// this is useful when u want to kill a goroutine and don't want to close the channels...

	time.Sleep(100 * time.Millisecond)
}
