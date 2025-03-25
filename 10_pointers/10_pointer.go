package main

import (
	"fmt"
)

func main() {

	var a int = 42
	var b *int = &a
	fmt.Println(a, b, &a, *b) // dereference a pointer * operator used

	a = 27
	fmt.Println(a, *b)

	*b = 14
	fmt.Println(a, *b)

	c := [3]int{1, 2, 3}
	d := &c[0]
	e := &c[1]
	fmt.Printf("%v %p %p\n", c, d, e)
	// e := &c[1] + 4 // not allowed - pointer arithmetic not is not present in go, if we want to use pointer arithmetic then we gotta use unsafe package

	type myStruct struct {
		foo int
	}

	var ms *myStruct
	ms = &myStruct{foo: 42}
	fmt.Println(ms)

	// new()
	ms = new(myStruct) // u can't initialization syntax
	fmt.Println(ms)
	(*ms).foo = 41 // (*ms) coz * operator has lower precedence than . operator
	fmt.Println(ms, (*ms).foo)

	// the compiler helps us by seeing (*ms).foo as ms.foo and hence we can omit (*)
	// the is a build in syntax in the complier, it understands that we want to access the struct underneath and not the address
	ms.foo = 26
	fmt.Println(ms, ms.foo)

	//* Zero value of the pointer is nil
	var ms2 *myStruct
	fmt.Println(ms2)

	// SLICES and MAPS copy works by copy the pointer and the data itself and hence any changes made to one reference will affect the other

	// PRIMITIVES, ARRAYS and STRUCTS copy the data itself (unless your using pointers)

}
