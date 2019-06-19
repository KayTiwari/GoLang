package main

import "fmt"

func main() {
	s := "H"

	bs := []byte(s)
	fmt.Println(bs[0])

	n := bs[0]
	fmt.Printf("%T\n", n)
	fmt.Printf("%b\n", n)
	fmt.Printf("%#x\n", n)
}

// numeral systems are systems used to keep track of numbers
// most numeral systems are base 10
// bits are in binary
// hexadecimal is in base of 16
