package main

import "fmt"

func splitSlice[T any](s []T) ([]T, []T){
	mid := len(s)/2
	return s[:mid], s[mid:]
}

func main(){
	firstInts, secondInts := splitSlice([]int {0,1,2,3,4,5,6,7,8,9})
	firstStrings, secondStrings := splitSlice([]string {"hello", "world"})
	fmt.Println(firstInts)
	fmt.Println(secondInts)
	fmt.Println(firstStrings)
	fmt.Println(secondStrings)
}