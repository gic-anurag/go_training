// Golang program to illustrate the usage of
// fmt.Fprint() function

// Including the main package
package main

// Importing fmt and os
import (
	"fmt"
	"os"
)

// Calling main
func main() {

	// Declaring some const variables
	const name, dept = "GeeksforGeeks", "CS"

	// Calling Fprint() function which returns
	// "n" as the number of bytes written and
	// "err" as any error ancountered
	n, err := fmt.Fprint(os.Stdout, name, " is a ",
		dept, " portal.\n")

	// Printing the number of bytes written
	fmt.Print(n, " bytes written.\n")

	// Printing if any error encountered
	fmt.Print(err)

}
