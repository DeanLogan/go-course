package main

import "fmt"

func main() {
	fmt.Println(oneReturnValue())
	firstReturnValue, _ := twoReturnValues()
	_, secondReturnValue := twoReturnValues()
	bothReturnValues1, bothReturnValues2 := twoReturnValues()
	fmt.Println(firstReturnValue)
	fmt.Println(secondReturnValue)
	fmt.Println(bothReturnValues1)
	fmt.Println(bothReturnValues2)
}

func oneReturnValue() string {
	return "Hello world!"
}

func twoReturnValues() (string, string) {
	return "Hello", "world!"
}

func initializedValues() (x, y int) {
	// x and y are initialized to 0
	return // automatically returns x and y
}
