// Create alias to long function names
package main

import (
	"fmt"
)

//Create alias to long function names
var pl = fmt.Println

// When a Go program executes it executes a function named main
// Go statements don't require semicolons
func main() {
	// ----- VARIABLES -----
	// var name type
	// Name must begin with letter and then letters or numbers
	// If a variable, function or type starts with a capital letter
	// it is considered exported and can be accessed outside the
	// package and otherwise is available only in the current package
	// Camal case is the default naming convention

	 var vName string = "Derek"
	 var v1, v2 = 1.2, 3.4

	// Short variable declaration (Type defined by data)
	 var v3 = "Hello"

	// Variables are mutable by default (Value can change as long
	// as the data type is the same)
	 v1 := 2.4

	// After declaring variables to assign values to them always use
	// = there after. If you use := you'll create a new variable

}
