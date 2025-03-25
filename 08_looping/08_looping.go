package main

import (
	"fmt"
)

func forLoop() {

	fmt.Println("\nFor Loop:")

	for i := 0; i < 5; i++ {
		fmt.Print(i)
	}
	fmt.Println()

	// for i := 0, j := 0; i < 5; i++, j++ { }  will not work

	// 1
	// can't have i := 0, j := 0
	// combine them i, j := 0, 0

	// 2
	// can't do i++, j++
	// i++ doesn't return anything, it is a statement and not an expression
	// so there no pre and post fix operators, here i++ is Increment statement and i-- is Decrement statement
	// combine them i, j = i+1, j+1

	for i, j := 0, 0; i < 5; i, j = i+1, j+2 {
		fmt.Print(i, j, ", ")
	}
	fmt.Println()

	// avoid messing with counter inside the loop
	for i := 0; i < 5; i++ {
		fmt.Print(i)
		if i%2 == 0 {
			i /= 2
		} else {
			i = 2*i + 1
		}
	}
	fmt.Println()

	i := 0
	for ; i < 5; i++ {
		fmt.Print(i)
	}
	fmt.Println(i)

	// while loop
	// while loop is nothing but for loop but rearranged (special case of for loop)
	i = 0
	for i < 5 {
		fmt.Print(i)
		i++
	}
	fmt.Println()

	// u can avoid the ; making it
	i = 0
	for i < 5 {
		fmt.Print(i)
		i++
	}
	fmt.Println()

	// do while loop
	i = 0
	for {
		fmt.Print(i)
		i++
		if i == 5 {
			break
		}
	}
	fmt.Println()

	//continue
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Print(i)
	}
	fmt.Println()

	// for in for
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Print(i * j)
		}
	}
	fmt.Println()

}

func labels() {

	fmt.Println("\nLabels:")

	// Label is used in break and continue statement where it’s optional but it’s required in goto statement.
	// Scope of the label is the function where it’s declared.

	// https://medium.com/golangspec/labels-in-go-4ffd81932339

	i := 0
start:
	fmt.Print(i)
	if i > 2 {
		goto end
	} else {
		i += 1
		goto start
	}
end:
	fmt.Println()

	// for goto, the label can be place be anywhere in the function

	// break and continue are used with for loop and switch statements
	// works for nested loops and nested switch and forloop
	// if we want to use labels with break and continue then we gotta write the label before any one of the outer for loops / before the switch statement

	// this is to avoid if statement after every if block in a for in for in for... u can jump to the first for or the 2nd for based on the label
outerLoop:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Print(i * j)
			if i*j >= 3 {
				break outerLoop
			}
		}
	}
	fmt.Println()

outerLoop2:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Print(i * j)
			if i*j >= 3 {
				continue outerLoop2
			}
		}
	}
	fmt.Println()

switchStatemnt:
	switch 1 {
	case 1:
		fmt.Print(1)
		for i := 0; i < 10; i++ {
			break switchStatemnt
		}
		fmt.Print(2)
	}
	fmt.Println(3)

	// 2 more rules

	// 1
	// any variable declaration cannot be skipped so causing a variable to come into scope if it wasn’t at the point where goto is used

	// 	goto Done
	//     v := 0
	// Done:
	//     fmt.Println(v)

	// throws a compiler error

	// 2
	// goto cannot move into other block
	// goto Block
	// {
	// Block:
	// 	v := 0
	// 	fmt.Println(v)
	// }

}

func forRangeLoop() {

	fmt.Println("\nFor Range Loop:")

	a := []int{1, 2, 3} // works with arrays or slices
	for k, v := range a {
		fmt.Println(k, v)
	}

	statePopulations := map[string]int{
		"California": 39250017,
		"Texas":      2750293,
		"Florida":    2932398,
		"New York":   29230929,
		"Ohio":       2039283092,
	}
	for k, v := range statePopulations {
		fmt.Println(k, v)
	}

	s := "Hello Go!"
	for k, v := range s {
		fmt.Println(k, string(v))
	}

	// only key
	for k := range statePopulations {
		fmt.Println(k)
	}

	// only value
	for _, v := range statePopulations {
		fmt.Println(v)
	}

}

func main() {

	forLoop()

	labels()

	forRangeLoop()

}
