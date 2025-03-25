package main

import (
	"fmt"
	"math"
)

func sayGreeting(greeting string, name string) {
	fmt.Println(greeting, name)
}

func sayGreeting2(greeting, name string) {
	fmt.Println(greeting, name)
}

func sayGreetingPassByValue(greeting, name string) {
	fmt.Println("\nPass By Value:")
	fmt.Println(greeting, name)
	name = "Ted"
	fmt.Println("Inside, name:", name)
}

func sayGreetingPassByReference(greeting, name *string) {
	fmt.Println("\nPass By Reference:")
	fmt.Println(*greeting, *name)
	*name = "Ted"
	fmt.Println("Inside, name:", *name)
}

// variadic function
// only one variadic parameter is allowed in a function and it has to be there at the end
func sum(msg string, values ...int) {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	fmt.Println(msg, result)

}

func sum2(values ...int) int {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	return result
}

// * return local variables
// In go language, when it (runtime) recoznizes that you're returning a value that's generated on the local memory (stack memory), it's automatically going to promote this variable for you to be on the shared memory (heap memory)
func sum3(values ...int) *int {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	return &result
}

// named return values
// not used much, be careful when choosing this
func sum4(values ...int) (result int) {
	fmt.Println(values)
	for _, v := range values {
		result += v
	}
	return
}

func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("Cannot divide by zero")
	}
	return a / b, nil
}

// methods

type greeter struct {
	greeting string
	name     string
}

func (g greeter) greet() {
	fmt.Println(g.greeting, g.name)
	g.name = "Deven2" // won't affect outside coz g here is a copy
}

func (g *greeter) greet2() {
	fmt.Println(g.greeting, g.name)
	g.name = "Deven2" // will affect outside coz g here is a pointer
}

type data int

func (d1 data) multiply(d2 data) data {
	return d1 * d2
}

func anonymousFunctionExamples() {

	fmt.Println("\nAnonymous Function Examples:")

	func() {
		msg := "Hello Go!" // msg can't be accessed outside this block / function
		fmt.Println(msg)
	}()

	for i := 0; i < 5; i++ {
		func() {
			fmt.Println(i) // can access variables from outside the function
			// but this might fail when we run this asynchronously
		}()
	}

	// solution to fix this for asynchronous
	// best practice coz changes in the outer scope aren't reflected in the inner scope
	for i := 0; i < 5; i++ {
		func(i int) {
			fmt.Println(i)
		}(i)
	}

}

func functionsAsTypesExamples() {

	fmt.Println("\nFunctions as Types:")

	f := func() {
		fmt.Println("Hello Go!")
	}
	f()

	var f2 func() = func() {
		fmt.Println("Hello Go!")
	}
	f2()

	var divide func(float64, float64) (float64, error)
	divide = func(a, b float64) (float64, error) {
		if b == 0.0 {
			return 0.0, fmt.Errorf("Cannot divide by zero")
		}
		return a / b, nil
	}
	res, err := divide(5.0, 0.0)
	if err != nil {
		fmt.Println(err)
		// return
	}
	fmt.Println(res)

	// we can't call this divde function before it is defined (assingment of the function) (at line 138)

}

func methodExamples() {

	// a method is just a function with a receiver argument.

	fmt.Println("\nMethod Examples:")

	// that variable is called the receiver for that method
	// the reciever can be value (copy is sent) or pointer (reference is sent)
	// same as functions with a different syntax to make it easier
	g := greeter{
		greeting: "Hello",
		name:     "Go",
	}
	g.greet()
	fmt.Println("The new name is:", g.name)

	g.greet2()
	fmt.Println("The new name is:", g.name)

	a := data(5)
	b := data(2)
	r := a.multiply(b)
	fmt.Println(r)

	// * You can only declare a method with a receiver whose type is defined in the same package as the method. You cannot declare a method with a receiver whose type is defined in another package (which includes the built-in types such as int).

	//  methods with pointer receivers take either a value or a pointer as the receiver when they are called:

	// var v Vertex
	// v.Scale(5)  // OK
	// p := &v
	// p.Scale(10) // OK

	// For the statement v.Scale(5), even though v is a value and not a pointer, the method with the pointer receiver is called automatically. That is, as a convenience, Go interprets the statement v.Scale(5) as (&v).Scale(5) since the Scale method has a pointer receiver. Vice Versa is also true (if the receiver is value type and pointer type is sent)

	// * Choosing a value or pointer receiver
	// There are two reasons to use a pointer receiver.

	// The first is so that the method can modify the value that its receiver points to.

	// The second is to avoid copying the value on each method call. This can be more efficient if the receiver is a large struct, for example.

	// In this example, both Scale and Abs are methods with receiver type *Vertex, even though the Abs method needn't modify its receiver.

	// * In general, all methods on a given type should have either value or pointer receivers, but not a mixture of both. (We'll see why over the next few pages.)

}

func variadicFunctionExamples() {

	fmt.Println("\nVariadic Function Examples:")

	sum("The sum is", 1, 2, 3, 4, 5)

	s := sum2(1, 2, 3, 4, 5, 6)
	fmt.Println("The sum is", s)

	s2 := sum3(1, 2, 3, 4)
	fmt.Println("The sum is", *s2)

	s3 := sum4(1, 2, 3)
	fmt.Println("The sum is", s3)

	d, err := divide(5.0, 0.0)
	if err != nil {
		fmt.Println(err)
		// return
	}
	fmt.Println(d)

}

func main() {

	sayGreeting("Hi", "Deven")
	sayGreeting2("Hello", "Deven")

	greeting := "Hi"
	name := "Deven"
	sayGreetingPassByValue(greeting, name)
	fmt.Println("Outside, name:", name)

	sayGreetingPassByReference(&greeting, &name)
	fmt.Println("Outside, name:", name)

	variadicFunctionExamples()

	anonymousFunctionExamples()

	//* functions as types
	functionsAsTypesExamples()

	// methods
	methodExamples()

	functionsAsParameters()

	functionClosures()

}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func functionsAsParameters() {

	// Functions are values too. They can be passed around just like other values.

	//Function values may be used as function arguments and return values.

	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func fibonacci() func() int {
	num1, num2 := 0, 1
	return func() int {
		prev := num1
		num1, num2 = num2, num1+num2
		return prev
	}
}

func functionClosures() {
	// Go functions may be closures. A closure is a function value that references variables from outside its body. The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.

	// For example, the adder function returns a closure. Each closure is bound to its own sum variable.

	// TODO: check this in python
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}

	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
