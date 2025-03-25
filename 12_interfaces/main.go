package main

import (
	"fmt"
	"io"
	"trial/bwc"
)

type Writer interface {
	Write([]byte) (int, error)
}

type ConsoleWriter struct {
}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

func interfaceForStructs() {

	// structs are data containers
	// interfaces don't describe data, interfaces describe behaviors, so we are storing method definitions

	// In go we don't explicitely implement interfaces, but we're going to implicitly implement the interface by creating a variable that's of type Writer and set that equal to a ConsoleWriter instance

	var w Writer = ConsoleWriter{}
	w.Write([]byte("Hello Go!"))

	// the w variable here is holding a Writer which is something that implements the Writer interface and it doesn't actually know the ConcreteWriter type though. So when the Write() method is called using w, it (w) knows how to call that (Write()) because that's defined by the interface BUT it (w.Write()) doesn't actually know the what's being written to, that's the responsibility of the actual implementation (ConsoleWriter()). So we could replace the ConsoleWriter with TCPWriter, FileWriter, etc. (this is a polymorphic behavior)

	// implicit implementation - if we need to wrap a concrete type and somebody hasn't published an interface, you can actually create an interface that their type implements.

	// So we don't have to worry about creating interfaces at design time if I don't need them myself because consumers of a library or whatever I'm creating can always create interfaces later and their interfaces CAN BE SHAPED TO EXACTLY WHAT THEY NEED FOR THEIR APPLICATION

	// eg. sql libraby has a concrete type (struct) and has a lot of methods with it. So if our application needs only 3 methods, i.e., connection(), writeAll(), readAll(), we make an interface with only these function signatures so that our application doesn't access other methods

	//* if we have a special case when we have a single method interface (which are very common in go language) then the convention is to name the interface with the - method name + "er" (eg. Write => Writer, Read => Reader)

	// for others name the interface by what it does

}

// interface using any type which has methods

type Incrementer interface {
	Increment() int
}

type IntCounter int

func (ic *IntCounter) Increment() int {
	*ic++
	return int(*ic)
}

func interfaceForOtherTypes() {

	// any type that can have a method associated with it can implement an interface

	myInt := IntCounter(0)
	var inc Incrementer = &myInt
	for i := 0; i < 10; i++ {
		fmt.Println(inc.Increment())
	}

	// primitive types can't be modified

}

type Closer interface {
	Close() error
}

type WriterCloser interface {
	Writer
	Closer
}

func composeInterfaceTogether() {

	var wc WriterCloser = bwc.NewBufferedWriterCloser()
	wc.Write([]byte("Hello Youtube listeners, this is a test"))
	wc.Close()

	//* Type conversion
	bwc := wc.(*bwc.BufferedWriterCloser) // look into type_assertions.go
	fmt.Println(bwc)

	// r := wc.(io.Reader)
	// fmt.Println(r)
	// with this syntax (only r returning) the application will panic as WriterCloser doesn't implement the Read method and io.Reader requires the Read method and the application doesn't know what to do so it panics

	// but with the below syntax we can avoid panics
	r, ok := wc.(io.Reader)
	if ok {
		fmt.Println(r)
	} else {
		fmt.Println("Conversion failed")
	}

}

func emptyInterface() {

	// empty interface has no methods and hence any type can be assigned to it even primitives

	// since it doesn't have methods it can't do anything useful until type conversion or using the reflect package in order to figure out what kind of an object ur dealing with

	var myObj interface{} = bwc.NewBufferedWriterCloser()
	if wc, ok := myObj.(WriterCloser); ok {

		wc.Write([]byte("Hello Youtube listeners, this is a test"))
		wc.Close()

	}

	r, ok := myObj.(io.Reader)
	if ok {
		fmt.Println(r)
	} else {
		fmt.Println("Conversion failed")
	}

}

func TypeSwitches() {

	// A type switch is a construct that permits several 'type assertions' in series.

	var i interface{} = 0
SwitchStatement:
	switch v := i.(type) {
	case int:
		fmt.Println("i is an integer", v)
		i = "0"
		goto SwitchStatement
	case string:
		fmt.Println("i is a string", v)
		i = 0.5
		goto SwitchStatement
	default:
		fmt.Println("I don't know what i is")
	}

	// instead of print statements the logic would be added of how to handle that type

}

func main() {

	// * An interface type is defined as a set of method signatures.

	// * A value of interface type can hold any value that implements those methods.

	// A type implements an interface by implementing its methods. There is no explicit declaration of intent, no "implements" keyword.

	// Implicit interfaces decouple the definition of an interface from its implementation, which could then appear in any package without prearrangement.

	interfaceForStructs()

	interfaceForOtherTypes()

	composeInterfaceTogether()

	emptyInterface()

	TypeSwitches()

	InterfaceValues()

	bestPractices()

}

func bestPractices() {

	// Best Practices

	// 1 Use many, small interfaces and compose them than using a huge monolitic one
	// useful when if a method requires only a Write function then we used use the Writer Interface as paramter, instead of using the WriterCloser interface and then pass the WriterCloser object to the Writer Interface paramter which will work just fine and that method can't accidently call the Close method coz it only knows about the Write method
	// eg. most powerful interface
	// io.Writer - 1 method, io.Reader - 1 method, interface{} - 0 method

	// 2 Don't export interfaces for types that will be consumed - struct in struct and the internal struct is being consumed by external sturct, so for the internal struct we should try to avoid exporting the interface, eg. the sql package the exported struct is made up of several structs and the internal structs interfaces are not exported

	// Do export interfaces for types that will be used by package - if we are going to pull a value in (like a paramter), go ahead and accept an interface instead of a concrete type, if at all possible - This is because of the implicit implementation

	// If we are creating a library that other people are going to consume, you can define the interfaces that you accept and they can provide whatever implementation that they want, if your library has resonable defaults then you could export those concrete types as well

	// 3 Design fucntions and methods to receive interfaces whenever possible - if we need access to the underlying data fields then we take in the concrete types. But if you're accepting behavior providers then accept those as interface types instead of concrete types

}
