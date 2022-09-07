package main

import "fmt"

// Repeat returns character repeated 5 times.
func Repeat(character string) string {
	var repeated string
	for i := 0; i < 5; i++ {
		//fmt.Println("\n")
		repeated = repeated + character
	}
	return repeated
}
func main() {
	fmt.Println("character repeation in a string")
	fmt.Println(Repeat("anurag"))
}
