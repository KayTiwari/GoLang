package main

//ints don't have decimals considered whole numbers
//floating point numbers do have decimals and are considered real numbers

import "fmt"

var x int
var y float64

func main() {
	x = 42
	y = 42.345
	//variable shadowing - when a variable declared in a certain scope has the same name 
	//as a variable in another scope

	//any number used in go is a constant

	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("%T\n", x)
	//use int for whole number no decimal points
	fmt.Printf("%T\n", y)
	//use float64 for when you need to use decimal points
}

//uint - unsigned: positive #s only including 0

//aliases
// byte is an alias for uint8
// rune is an alias for int32

//runtime package
//GOOS and GOARCH
