package main

import "fmt"

// const a = 42
// const b = 42.79
// const c = "James Bond"

const (
	a int = 42;
	b float64 = 42.79;
	c string = "James Bond"
)
//strict types or "typed"

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
