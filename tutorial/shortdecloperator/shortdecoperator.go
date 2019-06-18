// Identifiers name program entities such as types and variables.
// first char of identifiers must be a letter
// Extended Backus - Naur form

//Keywords may not be used as identifiers

// The short declarator operator is :=
package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)
	x = 99
	fmt.Println(x)
	y := 100 + 24;
	// this is a statement
	fmt.Println(y);
}
