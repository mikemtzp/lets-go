package main
// Variables
//  Can be defined by explicitly specifying a type.
var explicit int // Explicitly typed

// You can also use an initializer ':=', and the compiler will assign
// the variable type to match the type of the initializer
//  The return type(s) of a function must be specified in the function declaration.
func implicit() int {
  implicit := 10   // Implicitly typed as an int
  return implicit
}

// Once declared, a variable's type can never change.
func varType() int {
  count := 1 // Assign initial value
  count = 2  // Update to new value
  
  // count = false // This throws a compiler error due to assigning a non `int` type

  return count
}

// Constants
//  Their value cannot change during the execution of the program.
const Age = 21 // Defines a numeric constant 'Age' with the value of 21

// Functions
// Go functions accept 0 or more parameters.
// Parameters MUST BE explicitly typed, there is no type inference.
// Functions with 'PascalCase' will be public and with 'camelCase' will be private.

// Hello is a public function.
func Hello(name string) string {
  return hi(name)
}

// hi is a private function.
func hi(name string) string {
  return "hi " + name
}