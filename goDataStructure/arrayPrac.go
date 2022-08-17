package main

import "fmt"

func main() {
	fmt.Println("Arrays started")
	fmt.Println("")
	fmt.Println("arrays undefined size")
	m1()
	fmt.Println("")
	fmt.Println("to initialize the elements of specific index")
	m2()
	fmt.Println("")
	fmt.Println("slice creation")
	sliceCreationArr()
	fmt.Println("")
	fmt.Println("slice length")
	sliceLen()
	fmt.Println("")
	fmt.Println("access specific value of a map")
	accessMapVal()
	fmt.Println("")
	fmt.Println("accessing structure")
	stru()
	fmt.Println("")
	fmt.Println("access string elements")
	accessStr()
	fmt.Println("")
	fmt.Println("length of string")
	strLen()
	fmt.Println("")
	fmt.Println("concatinate string")
	concatStr()

}

//Declare an array of undefined size

func m1() {
	students := [...]string{"anurag", "santosh", "ramashankar"}
	fmt.Println(students)
}

// to initialize the elements of index 0 and 3 only
func m2() {

	arrayOfIntegers := [5]int{0: 7, 3: 9}

	fmt.Println(arrayOfIntegers)
}
