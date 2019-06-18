package main

import "fmt"

var y = 41

//var keyword allows you to declare a variable outside a function body

func main() {
	//short declaration operator - x, 42 is an operand, := is the operator
	//Declare a variable and assign a value at the exact same time
	x := 42
	fmt.Println(x)

	// var y = 43;
	fmt.Println(y)
}
