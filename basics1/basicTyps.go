package main

import (
	"fmt"
	"math/cmplx"
	"strings"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
	zeroVal()
	fmt.Println("")
	converterTool()
	fmt.Println("")
	aray()
	fmt.Println("")
	fmt.Println("Slices refrence to array")
	arrSlice()
	fmt.Println("")
	fmt.Println("slices of slice")
	slsl()
	fmt.Println("")
	fmt.Println("Append to a slice")
	slAppend()
	fmt.Println("")
	fmt.Println("Range use")
	rangeMeasure()
	fmt.Println("")
	fmt.Println("constants")
	consta()

}

func zeroVal() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}

//Conversion

func converterTool() {

	// taking the required
	// data into variables
	var totalsum int = 846
	var number int = 19
	var avg float32

	// explicit type conversion
	avg = float32(totalsum) / float32(number)

	// Displaying the result
	fmt.Printf("Average = %f\n", avg)
}

func aray() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

//Slices refrence to array

func arrSlice() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}

// slices of slice

func slsl() {
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

//Append to a slice

func slAppend() {
	var s []int
	printSlic(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlic(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlic(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlic(s)
}

func printSlic(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

//Range use

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func rangeMeasure() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

// constant
const Pi = 3.14

func consta() {
	const World = "welcome to go"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
