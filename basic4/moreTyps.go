package main

import "fmt"

func main() {
	fmt.Println("started")
	fmt.Println("pointer")
	point()
	fmt.Println("")
	fmt.Println("pointers to struct")
	pointToStruct()
	fmt.Println("")
	fmt.Println("slice")
	sliceLit()
	fmt.Println("")
	fmt.Println("length and capacity")
	lenCap()
	fmt.Println("")
	fmt.Println("slices by make keyword")
	makekey()

}

func point() {
	i, j := 42, 2701

	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)
}

// ponters to Struct

type Vertex struct {
	X int
	Y int
}

func pointToStruct() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}

func sliceLit() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}

//Slice length and capacity

func lenCap() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

//slices by make keyword

func makekey() {
	a := make([]int, 5)
	printSlices("a", a)

	b := make([]int, 0, 5)
	printSlices("b", b)

	c := b[:2]
	printSlices("c", c)

	d := c[2:5]
	printSlices("d", d)
}

func printSlices(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
