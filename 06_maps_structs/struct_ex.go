package main

import "fmt"

type A struct {
	B
	C
}

type B struct {
	i int
	j int
}

type C struct {
	i int
	k int
}

func main() {
	a := A{}
	a.B.i = 1
	a.j = 2
	a.k = 3
	fmt.Println(a)
}
