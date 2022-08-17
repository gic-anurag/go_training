// Program using range with string
package main

import "fmt"

func strRng() {

	s := "golang"

	for i, iteam := range s {
		fmt.Printf("%d= %c \n", i, iteam)
	}

}
