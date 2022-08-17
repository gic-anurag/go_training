// Program using range with map

package main

import "fmt"

func mapRange() {

	subjectmarks := map[string]int{"hindi": 10, "eng": 20, "math": 30}

	for subject, marks := range subjectmarks {
		fmt.Printf("%v, %v ", subject, marks)
	}
}
