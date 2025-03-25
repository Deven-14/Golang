package main

import (
	"fmt"
	"strings"
)

func arrayOperations() {

	fmt.Println("ArrayOperations:")

	// if we mention the size in the brackets or if we put '...' in the brackets then it is any array
	// and hence array has fixed size

	//* ARRAY IS VALUE-TYPE

	grades := [3]int{97, 85, 93}
	fmt.Printf("Grades: %v\n", grades)

	grades2 := [...]int{97, 85, 93}
	fmt.Printf("Grades2: %v\n", grades2)

	var students [3]string
	students[0] = "Lisa"
	students[1] = "Ahmed"
	fmt.Printf("Students: %v\n", students)
	fmt.Printf("Student #1: %v\n", students[1])
	fmt.Println("Number of Students: ", len(students))

	// array of arrays
	var identityMatrix [3][3]int = [3][3]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}
	fmt.Println("IdentityMatrix: ", identityMatrix)

	var identityMatrix2 [3][3]int
	identityMatrix2[0] = [3]int{1, 0, 0}
	identityMatrix2[1] = [3]int{0, 1, 0}
	identityMatrix2[2] = [3]int{0, 0, 1}
	fmt.Println("IdentityMatrix2: ", identityMatrix2)

	// copy / assignment
	//* ARRAY assignment does a deep copy
	a := [...]int{1, 2, 3}
	b := a //* deep copy
	b[1] = 5
	fmt.Println("a:", a)
	fmt.Println("b:", b)

	c := [2][2]int{{1, 0}, {0, 1}}
	d := c //* d is a deep copy of c
	d[1][0] = 5
	fmt.Println("c:", c)
	fmt.Println("d:", d)

	//* If we want to point to the same array then we gotta use pointers

	e := [...]int{1, 2, 3}
	f := &e // f is a pointer, pointing to array e
	f[1] = 5
	fmt.Println("e:", e)
	fmt.Println("f:", f)

	// [start:end), start inclusive, end exclusive
	//* all the ARRAYS created using an array will point to it and any changes made will reflect on the original array
	g := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	h := g[:]
	i := g[3:]
	j := g[:6]
	k := g[3:6]
	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)

	h[1] = 99
	i[1] = 69
	fmt.Println("g:", g) // g becomes {1, 99, 3, 4, 69, 6...}
	fmt.Println("h:", h)
	fmt.Println("i:", i)

}

func sliceOperations() {

	// A slice literal is like an array literal without the length.

	// This is an array literal:
	// [3]bool{true, true, false}

	// And this creates the same array as above, then builds a slice that references it:
	// []bool{true, true, false}

	// * The zero value of a slice is nil.

	fmt.Println("\nSliceOperations:")

	// when we don't mention the size in the brackets and leave it empty it becomes a slice
	// slice doesn't have fixed size and hence along with len() it has cap() [capacity]

	//* SLICE IS REFERENCE-TYPE

	a := []int{1, 2, 3}
	fmt.Println("a:", a)
	fmt.Println("Length:", len(a))
	fmt.Println("Capacity:", cap(a))

	//* slice points to the same array when assignment is performed
	// any changes to b will affect a
	b := a
	b[1] = 5
	fmt.Println("a:", a)
	fmt.Println("b:", b)

	// [start:end), start inclusive, end exclusive
	// all the SLICES created using a slice will point to it and any changes made will reflect on the original slice
	c := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	d := c[:]
	e := c[3:]
	f := c[:6]
	g := c[3:6]
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)

	d[1] = 99
	e[1] = 69
	fmt.Println("c:", c) // c becomes {1, 99, 3, 4, 69, 6...}
	fmt.Println("d:", d)
	fmt.Println("e:", e)

	// make function
	// when len == capacity a new slice is created with double the capacity, so a new slice with new capacity is created and values are copied to it and new values are added to it, when the slice gets big the copy operation can be expensive
	// so we can use the make function to give it a large enough capacity which is the min / max number of elements we know, which reduces the copy operation

	t := make([]int, 3) // makes a slice with 3 elements default value 0, here len = capacity = 3
	fmt.Println("t:", t)
	fmt.Println("Length:", len(t))
	fmt.Println("Capacity:", cap(t))

	i := make([]int, 3, 100) // makes a slice with 3 elements default value 0 (len - 3) and capacity 100
	fmt.Println("i:", i)
	fmt.Println("Length:", len(i))
	fmt.Println("Capacity:", cap(i))

	j := []int{}
	fmt.Println("j:", j)
	fmt.Println("Length:", len(j))
	fmt.Println("Capacity:", cap(j))

	// append
	j = append(j, 1)
	fmt.Println("j:", j)
	fmt.Println("Length:", len(j))
	fmt.Println("Capacity:", cap(j))

	j = append(j, 2, 3, 4, 5)
	fmt.Println("j:", j)
	fmt.Println("Length:", len(j))
	fmt.Println("Capacity:", cap(j))

	k := []int{6, 7, 8}
	j = append(j, k...)
	j = append(j, []int{9, 10, 11}...)
	fmt.Println("j:", j)
	fmt.Println("Length:", len(j))
	fmt.Println("Capacity:", cap(j))

	// pop operation
	j = j[1:]        // removing first ele
	j = j[:len(j)-1] // removing last ele
	index := 2       // removing index ele
	j = append(j[:index], j[index+1:]...)
	fmt.Println("j:", j)
	fmt.Println("Length:", len(j))
	fmt.Println("Capacity:", cap(j))

	// be careful of
	fmt.Println("before, j:", j)
	m := append(j[:index], j[index+1:]...)
	fmt.Println("after, m:", m)
	fmt.Println("after, j:", j)

	// before, j: [2 3 5 6 7 8 9 10]
	// after, m: [2 3 6 7 8 9 10]
	// after, j: [2 3 6 7 8 9 10 10]

	// slices of slices
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

}

func main() {

	arrayOperations()

	sliceOperations()

}
