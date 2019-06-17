package main

import "fmt"

func main() {
	//fmt.println returns bytes and err
	n, _ := fmt.Println("fuck you")
	//underscore ignores second return
	fmt.Println(n);
	foo()

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			// fmt.Println(i);
		}
	}
}

func foo() {
	fmt.Println("foo")
}
