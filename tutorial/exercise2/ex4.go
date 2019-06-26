//Write a program that:
//1. assigns an int to a variable
//2. prints that int to decimal, binary, and hex
//3. shifts the bits of that int over one position to the left, and assigns that to a variable
//4. prints that variable to decimal, binary, and hex

package main

import "fmt"

func main() {
	x := 666
	fmt.Printf("%d\t%b\t%#x\n", x, x, x)

	y := x << 1
	fmt.Printf("%d\t%b\t%#x\n", y, y, y)

}
