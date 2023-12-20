package main

import "fmt"

func main() {
	// for loops
	for i := 0; i < 5; i++ {
		fmt.Println("i:", i)
	}
	fmt.Println("====================================")
	// while loops
	j := 0
	for j < 5 {
		fmt.Println("j:", j)
		j++
	}
	fmt.Println("====================================")
	// infinite loops
	// for {
	// 	fmt.Println("infinite loop")
	// }
	// fmt.Println("====================================")
	// break
	for i := 0; i < 5; i++ {
		if i == 3 {
			break
		}
		fmt.Println("i:", i)
	}
	fmt.Println("====================================")
	// continue
	for i := 0; i < 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Println("i:", i)
	}
	fmt.Println("====================================")
	// nested loops
	for i := 0; i < 5; i++ {
		fmt.Println("i:", i)
		for j := 0; j < 3; j++ {
			fmt.Println("j:", j)
		}
	}
	fmt.Println("====================================")
	// labeled break
}