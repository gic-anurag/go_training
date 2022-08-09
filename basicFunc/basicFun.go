package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("started")
	fmt.Println("")
	fmt.Println("function with params")
	withParam(5, 7)
	fmt.Println("")
	fmt.Println("swaping")
	fmt.Println(swap(10, 30))
	fmt.Println("")
	fmt.Println("splint")
	split(45)
	fmt.Println(split(45))
	fmt.Println("")
	fmt.Println("Switch statement")
	Weakdays()
	fmt.Println("")
	fmt.Println("Differ")
	Dif()

}

func withParam(a, b int) int {
	fmt.Println(a + b)
	c := a + b
	fmt.Println(split(c))
	return c

}
func swap(x, y int) (int, int) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func Weakdays() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

func Dif() {
	defer fmt.Println("world")

	fmt.Println("hello")
}
