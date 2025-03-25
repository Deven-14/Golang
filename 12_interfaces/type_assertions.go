package main

import (
	"fmt"
	"io"
	"trial/bwc"
)

func typeAssertions() {

	// A type assertion provides access to an interface value's underlying concrete value.

	// t := i.(T)

	// This statement asserts that the interface value i holds the concrete type T and assigns the underlying T value to the variable t.

	// If i does not hold a T, the statement will trigger a panic.

	// To test whether an interface value holds a specific type, a type assertion can return two values: the underlying value and a boolean value that reports whether the assertion succeeded.

	// t, ok := i.(T)

	var wc WriterCloser = bwc.NewBufferedWriterCloser()
	wc.Write([]byte("Hello Youtube listeners, this is a test"))
	wc.Close()

	//* Type assertions / interface conversion
	bwc := wc.(*bwc.BufferedWriterCloser) // WriterCloser is type converted to BufferedWriterCloser, so now we can access the buffer directly with bwc but we can't access the buff with wc as it doesn't know of the data but only the functionality
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

	// bwc := wc.(*bwc.BufferedWriterCloser) // works
	// bwc := wc.(bwc.BufferedWriterCloser) // fails

	// this is because if we use a value type assignment then all the methods that implement the interface have to all have value receievers
	// wc := BufferedWriterCloser{} - value type assignment
	// all the methods that are in the interface should have a 'value type receiver'
	// if any one of the methods also has pointer type receiver then this throws an error

	// In short Method set of value is all methods which value receivers

	// but if some methods are pointer type and some are value type  OR  all methods are pointer type  OR all methods are value type then we can use pointer type assignment and it will work
	// wc := &BufferedWriterCloser{} - pointer type assignment

	// In short Method set of pointer is all methods, regardless of receiver type

}
