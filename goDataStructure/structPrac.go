package main

import "fmt"

func stru() {

	type Person struct {
		name string
		age  int
	}
	person1 := Person{"amit", 24}

	fmt.Println(person1)

	// access specifc element of struct

	fmt.Println(person1.age)
	// change value of structure field
	person1.age = 28
	fmt.Println(person1)
	fmt.Println(person1.age)

}
