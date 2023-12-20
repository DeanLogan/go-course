package main

import "fmt"

func main() {
	var myIntes [10]int // create an array with space for 10 ints
	fmt.Println(myIntes)
	primes := [6]int {2,3,5,7,11,13} // create an array with 6 predefined values
	fmt.Println(primes)
	sliceOfPrime := primes[1:4] // creates a slice of the primes array that will be an array of indexs 1-3 (index 4 is not included)
	fmt.Println(sliceOfPrime)
	slice := make([]int, 5, 10) // makes a new array of type int with length 5 and capacity 10 (capacity is the size of the underlying array). You can omitte capacity and capacity will default to the length
	fmt.Println(slice)
	fmt.Println(sum(1,2,3,4))
	slice = append(slice, 1) // append 1 to the end of slice
	fmt.Println(slice) 
	slice2d := [][]int{sliceOfPrime,slice} // 2d slice (slice of slices)
	fmt.Println(slice2d)

	// use of the range operator
	for i, num := range sliceOfPrime {
		fmt.Println(i, num)
	}
}

// vairadic function that allows the caller to have more flexability on how to call this function, this is recieved as a slice.
// we can do this by using the spread operator (...)
func sum(nums ...int) int{
	num := 0
	for i := 0; i<len(nums); i++ {
		num += nums[i] 
	}
	return num
}