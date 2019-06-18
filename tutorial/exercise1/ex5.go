package main

import "fmt"

type fuckyou int

var x fuckyou
var y int

func main() {
	y := int(x)
	fmt.Println(y)

}
