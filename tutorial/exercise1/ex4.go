package main

import "fmt"

type fuckyou int

var x fuckyou

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}
