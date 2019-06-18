package main

import "fmt"

var y string
func main(){
	//declare a variable to be a certain type
	//then assign a value of that type to the variable
	fmt.Printf("%T\n", y);
	var y = "Hello"
	fmt.Println(y);
}