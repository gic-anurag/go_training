package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("stared")
	fmt.Println("")
	fmt.Println("methods with pointrs")
	metsPoin()
	fmt.Println("")
	fmt.Println("Methods and pointer indirection")
	metsPoints()
	fmt.Println("")
	fmt.Println("ERRORS")
	err()
}

//Pointers and functions

type Vertex struct {
	X, Y float64
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func metsPoin() {
	v := Vertex{3, 4}
	Scale(&v, 10)
	fmt.Println(Abs(v))
}

// Methods and pointer indirection

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func metsPoints() {
	v := Vertex{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)

	p := &Vertex{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}

// ERRORS

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number : %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	next := float64(1)
	prev := float64(0)
	for math.Abs(next-prev) > .01 {
		prev, next = next, next-(next*next-x)/(2*next)
	}
	return next, nil
}

func err() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
