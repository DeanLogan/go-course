package main

import "fmt"

func main() {
	emptyMap := make(map[string]int) // creates empty map of strings and ints
	ages := map[string]int {
		"John":22,
		"Dean":21,
		"Bruce":54,
	}
	fmt.Println(emptyMap)
	fmt.Println(ages)
	ages["John"] = 32 // change an element
	fmt.Println(ages)
	elem := ages["John"] // get an element
	fmt.Println(elem)
	delete(ages, "Bruce") // delete a key from a map
	fmt.Println(ages)
	elem, found := ages["John"] // search for an element in the key (found is true as it is found)
	fmt.Println(elem, " ", found) 
	elem, found = ages["Bruce"] // search for an element in the key (found is false as the key no longer exists and elem is set to 0 as it was not found)
	fmt.Println(elem, " ", found) 
}