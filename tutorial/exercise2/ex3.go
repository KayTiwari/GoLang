//Create typed and untyped constants

package main

import "fmt"

const (
	a         = 42
	b         = 42.79
	c         = "Hello"
	d int     = 42
	e float64 = 56.89
	f string  = "Goodbye"
)

func main() {
	fmt.Println(a, b, c, d, e, f)
}
