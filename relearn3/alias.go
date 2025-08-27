package main

import "fmt"

// Standard type definition
type MyInt int

// Type alias
type AliasInt = int

func alias() {
	// var a MyInt = 10
	var b int = 20
	var c AliasInt = 30

	// This will not compile because MyInt and int are distinct types
	// fmt.Println(a == b)

	// This will compile because AliasInt and int are the same type
	fmt.Println(c == b) // Output: true
}
