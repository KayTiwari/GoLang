package main

import "fmt"

func main() {
	s := "Hello"
	fmt.Println(s)
	fmt.Println([]byte(s))
	//string -> slice of byte

	//converts to utf-8 numbers
	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U ", s[i])
	}
}
