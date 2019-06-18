package main

import "fmt"

var y = 42
var z string= "Broken clock"


func main() {

	fmt.Println(y)
	fmt.Printf("%T\n", y)

	// z = 43
	//can't reassign value types 
	fmt.Printf("%T\n", z)
}

//Go is a static programming language