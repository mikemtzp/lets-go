package main

import "fmt"

func main() {
	// single-line comments start with "//"
	// comments are just for documentation - they don't execute
	fmt.Println("Hello world!")

	var w int
	var x float64
	var y bool
	var z string

	// %v: to print the value of the argument in a default format.
	// %f: to print the argument as a floating-point number.
	// %t: for bolleans, prints the word true or false.
	// %q: a double-quoted string safely escaped with Go syntax.
	fmt.Printf("%v %f %t %q\n", w, x, y, z)
}