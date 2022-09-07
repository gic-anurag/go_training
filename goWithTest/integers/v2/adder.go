package main

import "fmt"

// Add takes two integers and returns the sum of them.
func Add(x, y int) int {
	return x + y
}
func main() {

	fmt.Println("addition")
	fmt.Println(Add(4, 8))
}
