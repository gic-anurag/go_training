package main

import "fmt"

func accessStr() {

	str := "golang"

	fmt.Printf("%c", str[0])
}

// leangth of string
func strLen() {
	str := "golang"

	i := len(str)
	fmt.Println(i)

}

// Program to concatenate two strings

func concatStr() {
	str := "i love"
	stk := "go programming"

	stkr := str + " " + stk

	fmt.Println(stkr)
}
