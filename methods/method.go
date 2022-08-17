package main

import "fmt"

func main() {

	d := no{5, 5}
	fmt.Println(d.sqrt())

	fmt.Println(regularFunc(d))

}

type no struct {
	a int
	b int
}

func (n no) sqrt() int {

	return n.a*n.a + n.b*n.b
}

func regularFunc(n no) int {
	return n.a*n.a + n.b*n.b
}
