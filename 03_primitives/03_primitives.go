package main

import (
	"fmt"
)

func boolOperations() {

	var n bool = true
	var v bool // default value is false
	fmt.Printf("%v, %T\n", n, n)
	fmt.Println(v)

	// * NOT AN ALIAS FOR OTHER TYPES (e.g. int)
	// Zero value is false, i.e., default value - declaration without initialization value is false

}

func integerOperations() {

	// integers

	// signed integers - int8 int16 int32 int64
	// int - on 32 bit system it is int32 and on 64 it is int64

	// unsigned integers = uint8 uint16 uint32 uint64
	// * unitptr
	// and uint

	// Zero value is 0, i.e., default value - declaration without initialization value is 0

	arithematicOperations()

	bitOperations()

	shiftOperations()

}

func arithematicOperations() {

	a := 10
	b := 3

	fmt.Println(3 / 2)
	// int / int = int
	fmt.Println(a+b, a-b, a*b, a/b, a%b)

	// ! int + int8 won't work, addition of 2 different type of int's won't work, u have to explicitly convert one of them

	// int / int = int
	// float / float = float

	// float / int and int / float won't work, the int or the float has to be explicitely converted to the other type

}

func bitOperations() {

	a := 10
	b := 3

	fmt.Println(a&b, a|b, a^b, a&^b)
	// &^ operator is both the bits are 0 then output 1 else 0

	// 10 => 1010
	//  3 => 0011
	// 10 &^ 3 = 0100

}

func shiftOperations() {

	a := 8

	fmt.Println(a << 3)
	fmt.Println(a >> 3)

}

func floatingPointOperations() {

	// floating point numbers
	// float32 float64
	// := for this syntax it is by default float64
	// supports arithematic operations and does not support bit and shift operations

	// Zero value is 0, i.e., default value - declaration without initialization value is 0

}

func complexNumbers() {

	// complex64 complex128

	var m complex64 = 1 + 2i
	fmt.Println(m)

	// supports arithematic operations and does not support bit and shift operations

	// Zero value is 0+0i, i.e., default value - declaration without initialization value is 0+0i

	fmt.Printf("%v, %T\n", real(m), real(m))
	fmt.Printf("%v, %T\n", imag(m), imag(m))
	var o complex64 = complex(5, 12)
	fmt.Printf("%v, %T\n", o, o)

}

func stringOperations() {

	// string
	s := "this is a string"
	fmt.Printf("%v, %v, %T\n", s[2], string(s[2]), s[2])

	// s[2] = "u" , not possible, strings are immutable

	// concatenation
	s2 := "this is also a string"
	fmt.Printf("%v, %T\n", s+s2, s+s2)

	// converting to bytes
	// used when moving around strings such as sending a string as a reply to a web service call or writing into a file etc
	// * byte is an alias for uint8
	b := []byte(s)
	fmt.Printf("%v, %T\n", b, b)

	// * now u can edit the string
	b[2] = 'U' // won't work with double quotes "U"
	fmt.Printf("%v, %T\n", b, b)

}

func runeOperations() {

	// rune - UTF-8 characters
	// * rune is a alias of int32
	// represented with single quotes ''

	r := 'a'
	var s rune = 'a'
	fmt.Printf("%v, %T\n", r, r)
	fmt.Printf("%v, %T\n", s, s)

}

func main() {

	boolOperations()

	integerOperations()

	floatingPointOperations()

	complexNumbers()

	stringOperations()

	runeOperations()

}
