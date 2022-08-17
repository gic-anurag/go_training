package main

import "fmt"

func m1() {
	// Program using range with array

	// array of numbers
	numbers := [5]int{21, 24, 27, 30, 33}

	// using range to iterate over the elements of array
	for index, item := range numbers {
		fmt.Printf("numbers[%d] = % \n", index, item)
	}

}
