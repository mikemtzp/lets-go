package main

import "fmt"

func main() {
	// Println formats using the default formats for its operands and writes to standard output.
	fmt.Println("Hello world!")

	var w int
	var x float64
	var y bool
	var z string

	// fmt.Printf - Prints a formatted string to standard out.
	// %v: to print the value of the argument in a default format.
	// %f: to print the argument as a floating-point number.
	// %t: for bolleans, prints the word true or false.
	// %q: a double-quoted string safely escaped with Go syntax.
	fmt.Printf("%v %f %t %q\n", w, x, y, z)
}

// fmt.Sprintf() - Returns the formatted string
func sprintf() (string, string) {
		// The %v variant prints the Go syntax representation of a value, it's a nice default.
	s := fmt.Sprintf("I am %v years old", 10)
	// I am 10 years old

	b := fmt.Sprintf("I am %v years old", "way too many")
	// I am way too many years old

	return s, b
}

