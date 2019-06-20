package main

import "fmt"

const a = 42
const b = 42.79
const c = "James Bond"

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n", a)
	//int
	fmt.Printf("%T\n", b)
	//float64
	fmt.Printf("%T\n", c)
	//string
}
