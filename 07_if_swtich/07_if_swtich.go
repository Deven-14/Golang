package main

import (
	"fmt"
	"math"
)

func ifStatement() {

	fmt.Println("\nIf Statement:")

	statePopulations := map[string]int{
		"California": 39250017,
		"Texas":      2750293,
		"Florida":    2932398,
		"New York":   29230929,
		"Ohio":       2039283092,
	}

	if true {
		fmt.Println("haha")
	}

	// if initializer; test { }
	if value, ok := statePopulations["Florida"]; ok {
		fmt.Println(value)
	}
	// ! fmt.Println(value)
	// won't work coz the scope of value is included in the above block

	if returnTrue() {
		fmt.Println("return true worked")
	}

}

func returnTrue() bool {
	fmt.Println("returning true")
	return true
}

func comparisonOperators() {

	fmt.Println("\nComparison Operators:")

	number := 50
	guess := 30

	if guess < number {
		fmt.Println("Too low")
	}
	if guess > number {
		fmt.Println("Too high")
	}
	if guess == number {
		fmt.Println("You got it!")
	}
	fmt.Println(number <= guess, number >= guess, number != guess)

}

func logicalOperators() {

	fmt.Println("\nLogical Operators:")

	guess := 30
	if guess < 1 || guess > 100 {
		fmt.Println("The guess must be greater than 1 and less than 100")
	}
	// short circuting - if the first comparsion guess < 1 is true then the others are not checked and the code block is executed

	if guess >= 1 && guess <= 100 {
		fmt.Println("good")
	}

	fmt.Println(!true, !false)

}

func ifElseStatement() {

	// Variables declared inside an if short statement are also available inside any of the else blocks.

	fmt.Println("\nIf Else Statement:")

	number := 50
	guess := 25

	if guess < 1 {
		fmt.Println("The guess must be greater than 1!")
	} else if guess > 100 {
		fmt.Println("The guess must be less than 100!")
	} else {
		fmt.Println("logic here", number, guess)
	}

}

func floatingPointComparison() {

	fmt.Println("\nFloating Point Comparisons:")

	myNum := 0.1

	fmt.Println(math.Pow(math.Sqrt(myNum), 2))
	if myNum == math.Pow(math.Sqrt(myNum), 2) {
		fmt.Println("These are same", myNum)
	} else {
		fmt.Println("These are different", myNum)
	}

	myNum = 0.123

	fmt.Println(math.Pow(math.Sqrt(myNum), 2))
	if myNum == math.Pow(math.Sqrt(myNum), 2) {
		fmt.Println("These are same", myNum)
	} else {
		fmt.Println("These are different", myNum)
	}

	// best method...
	// generate a error value and check if result is less than the error value
	myNum = 0.123456

	fmt.Println(math.Pow(math.Sqrt(myNum), 2))
	if math.Abs(myNum/math.Pow(math.Sqrt(myNum), 2)-1) < 0.001 {
		fmt.Println("These are same", myNum)
	} else {
		fmt.Println("These are different", myNum)
	}

}

func switchStatement() {

	fmt.Println("\nSwitch Statement:")

	switch 2 {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("not one or two")
	}

	// switch 2 {
	// case 1:
	// case 5:
	// case 10:
	// 	fmt.Println("1 | 5 | 10")
	// case 2:
	// case 4:
	// 	fmt.Println("2 | 4")
	// default:
	// 	fmt.Println("others")
	// }
	// this syntax is not there in go, instead we have the below syntax

	switch 4 {
	case 1, 5, 10:
		fmt.Println("one, five or ten")
	case 2, 4, 6:
		fmt.Println("two, four or six")
	default:
		fmt.Println("another number")
	}
	// we cannot have a same value in 2 cases, the generates a syntax error

	switch i := 2 + 3; i {
	case 1, 5, 10:
		fmt.Println("one, five or ten")
	case 2, 4, 6:
		fmt.Println("two, four or six")
	default:
		fmt.Println("another number")
	}

	i := 10
	switch {
	case i <= 10:
		fmt.Println("less than or equal to ten")
	case i <= 20:
		fmt.Println("less than or equal to twenty")
	default:
		fmt.Println("greater than twenty")
	}

	// fallthrough can be used only as a last statement inside case clause
	switch {
	case i <= 10:
		fmt.Println("less than or equal to ten")
		fallthrough // continues even if the test case passes
	case i <= 20:
		fmt.Println("less than or equal to twenty")
	default:
		fmt.Println("greater than twenty")
	}

	// Type Switch

	var j interface{} = "abc" // 1, 1.32, "ab", [3]int{}, [2]int{}
	switch j.(type) {
	case int:
		fmt.Println("j is an int")
		fmt.Println("This will run too")
	case float64:
		fmt.Println("j is a float64")
		break // commented coz of warning unreachable code
		fmt.Println("This won't run")
	case string:
		fmt.Println("j is a string")
	case [3]int: // [3]int is different than [2]int
		fmt.Println("j is [3]int")
	case [2]int:
		fmt.Println("j is [2]int")
	default:
		fmt.Println("j is another type")
	}

}

func main() {

	ifStatement()

	comparisonOperators()

	logicalOperators()

	ifElseStatement()

	floatingPointComparison()

	switchStatement()

}
