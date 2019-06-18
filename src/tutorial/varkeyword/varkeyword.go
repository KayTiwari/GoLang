package main

import "fmt"

//var keyword allows you to declare a variable outside a function body
var y = 41

//declares a variable with the identifier z of the type int
// default value of type int is 0
var z int

func main() {
	//short declaration operator - x, 42 is an operand, := is the operator
	//Declare a variable and assign a value at the exact same time
	x := 42
	fmt.Println(x)

	// var y = 43;
	fmt.Println(y)
}
