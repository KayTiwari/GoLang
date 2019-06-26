//Using iota, create constants for the past 4 years

package main

import "fmt"

const (
	a = 2015 + iota
	b  = 2015 + iota
	c  = 2015 + iota
	d  = 2015 + iota
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
