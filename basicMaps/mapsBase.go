package main

import (
	"fmt"
	"math"
	"strings"

	"golang.org/x/tour/wc"
)

func main() {
	fmt.Println("Started")
	mapVert()
	fmt.Println("")
	fmt.Println("Mutetating map")
	muteMap()
	fmt.Println("")
	fmt.Println(" Fibbo series exersice")
	fibbo()
	fmt.Println("")
	fmt.Println("Word count")
	Exercise()
	fmt.Println("")
	fmt.Println("Functions are values too")
	FunRval()

}

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func mapVert() {
	fmt.Println(m)
}

// Mutetating map

func muteMap() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}

// Fibbo series exersice

func fibonacci() func() int {
	current := 1
	last := -1
	return func() int {
		current, last = (last + current), current
		return current
	}
}

func fibbo() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

//Word count

func WordCount(s string) map[string]int {
	//using make to initialise an empty map
	stringSlice := make(map[string]int)

	//words is a slice of substrings of s
	words := strings.Fields(s)

	for _, v := range words {
		//If key is not in the map,
		//then elem is the zero value for the map's element type.
		//zero value of int is 0
		//incrementing it by 1 for the number of times the key is found
		stringSlice[v]++

	}
	return stringSlice
}
func Exercise() {
	wc.Test(WordCount)
}

// Functions are values too

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func FunRval() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}
