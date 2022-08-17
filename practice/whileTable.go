package main

import "fmt"

func main() {
	fmt.Println("")
	fmt.Println("multiplication table with while")
	mulWithWhile()
	fmt.Println("")
	fmt.Println("Range with array")
	m1()
	fmt.Println("")
	fmt.Println("string with range")
	strRng()
	fmt.Println("")
	fmt.Println("Map range")
	mapRange()
	fmt.Println("")
	fmt.Println("get map keys")
	getKeyOfMap()
}

// multiplication table with while
func mulWithWhile() {
	multiplier := 1

	// run while loop for 10 times
	for multiplier <= 10 {

		// find the product
		product := 5 * multiplier

		// print the multiplication table in format 5 * 1 = 5
		fmt.Printf("5 * %d = %d\n", multiplier, product)
		multiplier++
	}
}
