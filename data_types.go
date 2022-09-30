// Create alias to long function names
package main

import (
	"fmt"
	"reflect"
)

//Create alias to long function names
var pl = fmt.Println

// When a Go program executes it executes a function named main
// Go statements don't require semicolons
func main() {
	


// ----- DATA TYPES -----
	// int, float64, bool, string, rune
	// Default type 0, 0.0, false, ""
	pl(reflect.TypeOf(25))
	pl(reflect.TypeOf(3.14))
	pl(reflect.TypeOf(true))
	pl(reflect.TypeOf("Hello"))
	pl(reflect.TypeOf('🦍'))
}