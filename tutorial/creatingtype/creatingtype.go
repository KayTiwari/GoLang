package main

import "fmt"

var a int = 42

type hotdog int

var b hotdog

func main() {
	a = 42
	b = 47
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	
	// a = b
	//can't redeclare a to type hotdog
	
	a = int(b)
	fmt.Println(a)
	fmt.Printf("%T\n", a)
}

//conversion - converting values of a certain type to another type
