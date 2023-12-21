package main

import "fmt"

func main() {
	defer stoppingFunc() // the defer keyword allows you to execute a function whenever the current function has ended

	squareFunc := selfMath(multiply) // converts the multiply function into a function that squares a number
	doubleFunc := selfMath(add) // converts the add function into a function that doubles the given number

	fmt.Println(squareFunc(5))
	fmt.Println(doubleFunc(5))

	harryPotterAggregator := concatter()
	harryPotterAggregator("Mr.")
	harryPotterAggregator("and")
	harryPotterAggregator("Mrs.")
	harryPotterAggregator("Dursley")
	harryPotterAggregator("of")
	harryPotterAggregator("number")
	harryPotterAggregator("four,")
	harryPotterAggregator("Privet")

	fmt.Println(harryPotterAggregator("Drive"))
}

func multiply(x, y int) int {
	return x * y
}

func add(x, y int) int {
	return x + y
}

// currying function, this is when a function takes a function(s) as input and returns a new function
func selfMath(mathFunc func(int, int) int) func (int) int {
	return func(x int) int {
		return mathFunc(x, x)
	}
}

func stoppingFunc(){
	fmt.Println("This function was defered and is happening at the end of this function")
}

// A closure is a function that references variables from outside its own function body. The function may access and assign to the referenced variables.
func concatter() func(string) string {
	doc := ""
	return func(word string) string {
		doc += word + " "
		return doc
	}
}