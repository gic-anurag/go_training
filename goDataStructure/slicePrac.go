package main

import "fmt"

func sliceCreationArr() {
	arr := [6]int{1, 2, 3, 4, 5, 6}

	sli := arr[2:4]
	fmt.Println(sli)

	sl := arr[:6]
	fmt.Println(sl)
}

func sliceLen() {

	slL := []int{1, 2, 3, 4, 5, 6, 7, 8}

	no := len(slL)
	fmt.Println(slL)
	fmt.Println(no)

	// iterate throgh the slice

	for i := 0; i < len(slL); i++ {
		fmt.Println(slL[i])
	}

}
