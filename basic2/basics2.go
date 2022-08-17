package main

import (
	"fmt"
)

func main() {
	fmt.Println("started")
	forState()
	fmt.Println("")
	fmt.Println("For loop without initialision and post statement")
	noInitPost()
	fmt.Println("")
	fmt.Println("for infinity")
	infin()
	fmt.Println("")
	fmt.Println("if else prac")
	ifTest()

}

func forState() {

	var count int = 10
	for i := 0; i < count; i++ {
		fmt.Println(i)
	}
}

// For loop without initialision and post statement
func noInitPost() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

func infin() {}

func ifTest() string {

	if a := 10; a < 50 {
		fmt.Println("A is smaller then 50")
		a := "true"
		return a
	} else {

		fmt.Println("A is greater then 50")
		a := "false"
		return a
	}

}
