package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
}

type Doctor struct {
	name string
	age  int
}

type data int

// we are creating new types from existing types such as struct and int
