package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func deferExample1() {

	fmt.Println("\nDefer Example 1:")
	fmt.Println("start")
	defer fmt.Println("middle")
	fmt.Println("end")

}

func deferExample2() {

	fmt.Println("\nDefer Example 2:")
	defer fmt.Println("start")
	defer fmt.Println("middle")
	defer fmt.Println("end")

	// output - end, middle, start
	// defer works in reverse order

}

func deferExample3() {

	fmt.Println("\nDefer Example 3:")

	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	robots, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", robots[:50]) // sliced the robots strings so the output terminal is clean

}

func deferExample4() {

	fmt.Println("\nDefer Example 4:")
	a := "start"
	defer fmt.Println("a:", a)
	a = "end"

	// ouput is - start
	// call function takes the arguments at the time the defer is called and not at the time the call function is executed

}

func deferExamples() {

	// defer - to delay the exuecution of that particular statemnet until the execution of the last statement of that function but before the return of the function

	// helps to group the open and close of the resource so that we don't forget it later on

	// if we are looping over many resources than it is recomended to close them manually than using defer

	// defer statments follow LIFO (last-in, first-out) order

	// call function takes the arguments at the time the defer is called and not at the time the call function is executeds

	deferExample1()

	deferExample2()

	deferExample3()

	deferExample4()

}

func panicExample1() {

	fmt.Println("\nPanic Example 1:")

	a, b := 1, 0
	ans := a / b
	fmt.Println(ans, a, b)

}

func panicExample2() {

	fmt.Println("\nPanic Example 2:")

	fmt.Println("start")
	panic("something bad happend")
	fmt.Println("end")

}

func panicExample3() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Go!"))
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err.Error())
	}

	// we are the ones writing web application and we know that if that (ListenAndServe()) doesn't work, the entire application gets brought down and nothing happens. So in that situation, we decide that we're going to panic (if err != nil { panic() })

}

func panicExamples() {

	fmt.Println("\nPanic Examples:")

	// panic - when the application can't continue to function so it's really starting to panic because it can't figure out what to do

	// opening a file which does not exist is not considered as exception. It is reasonable to assume you might try and open a file that doesn't exist and so we return error values and we don't throw errors as that is not considered exception in go

	// based on the error values we decide if we want to go into panic or not

	// however there are situations where a go application runs into a situation it cannot continue and that is considered exceptional. (like 1 / 0)

	//* go is rarely going to set an opnion about whether an error is something that should be panicked over or not, all it's going to do is tell you, is hey this didn't work the way you expected it to work. It's up to you as a developer to decide whether that's a problem or not. (If problem then panic else don't)

	//* uncomment any one function it see how the application panics

	// panicExample1()

	// panicExample2()

	// panicExample3()
	//* run this 09_defer_panic_recover.go program twice to see the panic as the port is already taken

	// so in this case we say that there is a problem, we fail to start our application and so we do want to generate a panic

	// panic don't have to be fatal. If they are, then we panic all the way up to the go runtime and then the go runtime realizes it doesn't know what to do with a panicking application and so it's going to kill it.

}

func panicAndDeferCombinedExample() {
	fmt.Println("\nPanic and Defer Combined Example:")
	fmt.Println("start")
	defer fmt.Println("this was deferred")
	panic("somthing bad happened")
	fmt.Println("end")

	// defer is performed first if the program goes into panic and then the panic comes into picture

	// so even if the program panics it will first close resources which are defered and then go into panic
}

func recoverExample1() {
	fmt.Println("\nRecover Example 1:")
	fmt.Println("start")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error:", err)
		}
	}()
	panic("something went bad")
	fmt.Println("end")

	// even after recover, the statements after panic() are not executed coz it assumes that if this function went into panic then whatever ahead of panic is there it shouldn't work coz of the panic
}

func panicker() {
	fmt.Println("about to panic")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error:", err)
			// instead of logging we'll write statement to recover from the error, if we can't then we panic again
		}
	}()
	panic("something went bad")
	fmt.Println("done panicking")
}

func recoverExample2() {
	fmt.Println("\nRecover Example 2:")
	fmt.Println("start")
	panicker()
	fmt.Println("end")

	// the statments after the panic() in the panicker() function are not executed coz that function should be unfit to perform now that it panicked but the statements after the panicker() in this function are executed as we recovered from that panic() in that function and so this function should be fit to execute
}

func panicker2() {
	fmt.Println("about to panic")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error:", err)
			panic(err) // we are panicking again coz here we assume we don't know how to recover from this err
			// if we know how to recover from the err we would perform operations to recover here and not panic()
		}
	}()
	panic("something went bad")
	fmt.Println("done panicking")
}

func recoverExample3() {
	fmt.Println("\nRecover Example 3:")
	fmt.Println("start")
	panicker2()
	fmt.Println("end")

	// the statments after the panic() in the panicker2() function are not executed coz that function should be unfit to perform now that it panicked and also the statements after the panicker2() in this function are  not executed as we panic() again after the recover() if we don't know how to handle it. So if we don't recover it again then it will reach the runtime.
}

func recoverExamples() {

	// used to recover from a panic
	// only useful in deferred functions
	// current function will not attempt to continue, but higher functions in call stack will

	fmt.Println("\nRecover Examples:")

	// recoverExample1()

	// recoverExample2()

	// recoverExample3()

	fmt.Println("Recover Examples done")

}

func main() {

	deferExamples()

	panicExamples()

	// panicAndDeferCombinedExample()

	recoverExamples()

}
