package main

import "fmt"

func defX() {

	for i := 0; i < 6; i++ {
		defer fmt.Println(i)
	}
}
