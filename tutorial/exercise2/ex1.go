//write a program that prints a number in decimal, binary, and hex
package main

import "fmt"

func main() {
	x := 666
	fmt.Printf("%d\t\t%b\t\t%#x\n", x, x, x)
}
