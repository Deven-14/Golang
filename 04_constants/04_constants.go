package main

import (
	"fmt"
)

// upper case first letter - PascalCase - if we want to export the const variable
// lower case first letter - camelCase - for package scope

const e int16 = 27

func constOperations() {

	const myConst int = 42
	fmt.Printf("%v, %T\n", myConst, myConst)

	// myConst = 5 // can't modify constant
	// myConst = math.Sin(1.57) // can't do this coz the value should be replacable at compile time, math.Sin() will require execution during runtime

	const a int = 15
	const b float32 = 3.14
	const c string = "hello"
	const d bool = true
	fmt.Println(a, b, c, d)

	// constants can also be declared again in different scopes, global and main funciton block here
	// * this is shadowing, so constants can shadowed
	const e int16 = 14
	fmt.Printf("%v, %T\n", e, e)

	// we can perform operations on constant and variables
	const f int = 5
	var g int = 6
	fmt.Printf("%v, %T\n", f+g, f+g)

	// another way to declare and initialize
	const h = 42

	// the compiler replaces the symbol 'h' everywhere in the code with the value 42 and hence the below is possible
	var i int16 = 27
	fmt.Printf("%v, %T\n", h+i, h+i) // seen as 42 + i
	// the resulting type is a int16
	// this is not possible through var but is possible through const coz of the replacing effect

}

const (
	_ = iota
	catSpecialist
	dogSpecialist
	snakeSpecialist
)

func enumeratedConstantsOperations() {

	fmt.Println("\nEnumerated Constants :")

	const a = iota
	const b = iota
	// iota special symbol that is a counter used when we are creating enumerated constants

	// iota is scoped in a const block and for each scope iota starts from 0

	const (
		c = iota
		d = iota
		e = iota
	)

	const (
		f = iota
		g
		h
	)

	const i = iota

	fmt.Println(a, b, i)
	fmt.Println(c, d, e)
	fmt.Println(f, g, h)

	// * MAINLY used in the package scope (basicaly like ENUM) eg. the specialist const scope

	var specialistType int = catSpecialist
	fmt.Println(specialistType == catSpecialist)

	// now if we don't assign a value to the specialistType then by default it is 0 and this will direct to catSpecialist, to avoid this error we can assign the first iota to errorSpecialist or _
	// * '_' is a write only variable
	// errorSpecialist can be used if we know that the specialistType can reach a value 0
	// if the program is built in such a way that it never reaches 0 then we can use _ the write only variable

}

const (
	isAdmin = 1 << iota
	isHeadquaters
	canSeeFinancials

	canSeeAfrica
	canSeeAsia
	canSeeEurope
	canSeeNorthAmerica
	canSeeSouthAmerica
)

func enumeratedExpressionsOperations() {

	// enumerated expressions - operations that can be determined at compile time are allowed
	// iota can be used with operations - arithematic, shift, bit

	// * compiler uses the same syntax downwards and iota is incremented for each const as we go down
	const (
		_ = iota + 5 // here iota will be 0, so 0 + 5 = 5
		j            // j = iota + 5 and iota = 1, so 1 + 5 = 5
		k
		l
	)
	fmt.Println(j, k, l)

	const (
		_  = iota
		KB = 1 << (10 * iota)
		MB
		GB
		TB
	)
	fmt.Println(KB, MB, GB, TB)

	fileSize := 4000000000.
	fmt.Printf("%.2fGB", fileSize/GB)

	// how the open function flags work
	// used byte so only 8 values should be there in the const iota scope
	var roles byte = isAdmin | canSeeFinancials | canSeeEurope
	fmt.Printf("%b\n", roles)
	fmt.Printf("Is Admin? %v\n", isAdmin&roles == isAdmin)
	fmt.Printf("Is HQ? %v\n", isHeadquaters&roles == isHeadquaters)

}

func main() {

	// * cannot be declared with := syntax

	// * An untyped constant takes the type needed by its context.
	const (
		a = 1 << 100
		b = a >> 99
	)

	constOperations()

	enumeratedConstantsOperations()

	enumeratedExpressionsOperations()

}
