package main

import "fmt"

func main() {
	var hello string = "hello" // Declare a variable named hello which stores a string
	world := "world"           // Declare a variable named world which stores a string using short variable declaration and assign it the value 0
	
	const firstName = "Dean"
	const lastName = "Logan"
	const fullName = firstName + " " + lastName // const variables are computed at compile time not runtime like other variables
	fmt.Println(hello, world)
	fmt.Println(fullName)
	formattingStrings()
}

func formattingStrings() {
	/*
		%v - Interpolate the default representation
		The %v variant prints the Go syntax representation of a value. You can usually use this if you're unsure what 
		else to use. That said, it's better to use the type-specific variant if you can.
	*/
	fmt.Printf("I am %v years old", 10)
	// I am 10 years old
	
	fmt.Printf("I am %v years old", "way too many")
	// I am way too many years old
	
	/*
		%s - Interpolate a string
	*/
	fmt.Printf("I am %s years old", "way too many")
	// I am way too many years old
	
	/*
		%d - Interpolate an integer in decimal form
	*/
	fmt.Printf("I am %d years old", 10)
	// I am 10 years old
	
	/*
		%f - Interpolate a decimal
	*/
	fmt.Printf("I am %f years old", 10.523)
	// I am 10.523000 years old
	
	// The ".2" rounds the number to 2 decimal places
	fmt.Printf("I am %.2f years old", 10.523)
	// I am 10.53 years old
}
