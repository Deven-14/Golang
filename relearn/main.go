package main

import "fmt"

func main() {

	// relearn1()

	// relearn2()

	// if err := runError(); err != nil {
	// 	fmt.Println(err)
	// }

	// runErrors2()

	// reader()

	// reader2()

	// image_print()

	// typeParameters()

	// goroutinues()

	// expression := "+ * 2 3 4"
	// result := Solve(expression)
	// fmt.Println(result) // Output: 10

	// expression2 := "- 10 / 2 3"
	// result2 := Solve(expression2)
	// fmt.Println(result2) // Output: 10

	input := `6 7
	1 2
	1 3
	2 4
	2 5
	3 4
	4 6
	5 6`
	result := Solve3(input)
	fmt.Println(result)

}
