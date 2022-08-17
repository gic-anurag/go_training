// Program to retrieve the keys of a map

package main

import "fmt"

func getKeyOfMap() {

	subjectMarks := map[string]int{"math": 20, "science": 50, "S.st": 80}

	for subject := range subjectMarks {
		fmt.Println(subject)
	}

}
