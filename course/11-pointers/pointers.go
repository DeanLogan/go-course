package main

import "fmt"

func main(){
	x := 6
	var p *int = &x // points to the memory address of the variable x, as x is an int, int needs to be used to define the pointer
	fmt.Println(p) // prints the memory address
	fmt.Println(*p) // prints the contents at this memory address (de-reference a pointer)
	// pointers allow for pass by reference instead of pass by value. This allows a change in a function to affect the variable everywhere
}