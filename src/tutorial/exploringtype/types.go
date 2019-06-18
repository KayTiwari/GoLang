package main

import "fmt"

var y = 42
var z string = "Broken clock"

var a string = `they said "hello"`
//backquotes allow you to put double quotes in a raw string literal

func main() {

	fmt.Println(y)
	fmt.Printf("%T\n", y)

	// z = 43
	//can't reassign value types
	fmt.Printf("%T\n", z)

	fmt.Println(a)
}

//Go is a static programming language

//slice data type- many values of one type
//struct - many values of different data types