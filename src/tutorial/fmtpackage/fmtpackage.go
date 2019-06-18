package main

import "fmt"

var y = 42

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	//binary
	fmt.Printf("%b\n", y)
	//hex
	fmt.Printf("%x\n", y)
}

//fmt package is similar to C's printf and scanf
