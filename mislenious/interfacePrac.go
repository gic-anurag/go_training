// Program to pass pointer as a function argument

package main

import "fmt"

func main() {

	rec := Rectangle{10, 10}
	fmt.Println("started")
	calc(rec)
	fmt.Println("")
	fmt.Println("generic practice")
	genPrac(10)
	//fmt.Println("")
	//fmt.Println("goroutine")
	//	go say("world")
	//	say("hello")
	fmt.Println("")
	//fmt.Println("goroutine sum")
	//go sum(10, 20)
	fmt.Println("")
	fmt.Println("defer")
	defer fmt.Println("defer")
	defer fmt.Println("example")
	defX()

}

func sum(a, b int) {
	c := a + b
	fmt.Println(c)

}

//Generic type

var age1 int = 10
var age2 float32 = 29.6

func genPrac[age any](myage age) {
	fmt.Println(myage)
}

type shape interface {
	area() int
}

func (r Rectangle) area() int {
	return r.length * r.breath
}

type Rectangle struct {
	length int
	breath int
}

func calc(s shape) {
	fmt.Println(s.area())
}
