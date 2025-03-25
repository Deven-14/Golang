package main

import (
	"fmt"
	"strconv"
)

// * at package level we cannot use := syntax
var l string = "abc"

// this syntax helps group related variables
var (
	actorName string = "Daniel RAdcliffe"
	movieName string = "Harry Potter"
)

var (
	nbooks   int = 5
	npens    int = 10
	npencils int = 15
)

var m int = 10

// variables starting with a upper case letter are exported, the rest can be accessed only inside the package
var I int = 100

// * 3 scopes of visiblity in Go
// UPPER CASE first letter at package level - exported and globaly accessible // * PascalCase
// lower case first letter at package level - only accessible in the package // * camelCase
// block scope - only accessible in that block (no private scope)

func main() {

	// 4 ways to declare and initialize
	var i int          // declare // * helpful when we want a variable to be declared in another scope
	i = 42             // initialize
	var j float32 = 27 // declare and initialize
	var z = 53         // declare and initialize
	k := 99.           // declare and initialize // mostly used

	// multiple
	var t, u int = 1, 2
	var w, x, y = 5, "oh", 3.5

	fmt.Println(i, j, z, k, l, m)
	fmt.Println(t, u, w, x, y)
	fmt.Printf("%v, %T\n", j, j)
	fmt.Printf("%v, %T\n", k, k)

	// GLOBAL AND BLOCK SCOPE
	// var m with inner most scope is used, u can declare the same variable in different scopes
	fmt.Println(m)
	var m int = 5
	fmt.Println(m)
	// m := 4 // cannot do this coz m is already declared in the above line but we can assign
	m = 3
	fmt.Println(m)

	// the ACRONYMS should be upper case
	var theURL string = "http://google.com"
	var theHTTPRequest string = "http://google.com"
	fmt.Println(theURL, theHTTPRequest)
	// * writeDB  upper case acronym
	// * dbWrite  lower case acronym and then upper case W

	// TYPE CONVERSIONS
	// go doesn't have implicit type conversions
	var n int = 67
	var o float32 = float32(n)
	p := 67.5
	var q int = int(p)
	fmt.Printf("%v, %T\n", o, o)
	fmt.Printf("%v, %T\n", q, q)

	var r string = string(rune(n))
	// rune is int32 (hover for explanation)
	// converts to UTF-8 representation of that integer
	var s string = strconv.Itoa(q)
	fmt.Printf("%v, %T\n", r, r)
	fmt.Printf("%v, %T\n", s, s)

}
