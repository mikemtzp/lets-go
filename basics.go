package main

// Variables
//  Can be defined by explicitly specifying a type.
var explicit int // Explicitly typed

// You can also use an initializer ':=', and the compiler will assign
// the variable type to match the type of the initializer
//  The return type(s) of a function must be specified in the function declaration (int).
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

// Type sizes
// whole numbers: int  int8  int16  int32  int64 
// positive whole numbers: uint uint8 uint16 uint32 uint64 uintptr 
// decimal numbers: float32 float64 
// imaginary numbers (rare): complex64 complex128

// Types can be converted the following way
func typeConvertion() int {
  f := 88.26
  n := int(f)

  return n
}

// Type inference: the variable's type is inferred from the value on the right hand side.
func typeInference() (int, int, float64, complex128) {
  var a int
  b := a // j is also an int

  // However, when the right hand side is a literal value (an untyped numeric constant like 42 or 3.14),
  // the new variable will be an int, float64, or complex128 depending on its precision:
  i := 42           // int
  f := 3.14         // float64
  c := 0.867 + 0.5i // complex128

  return b, i, f, c
}

// Same line declarations: we can declare multiple variables on the same line:
func sameLineDecl() (int, string) {
  order, command := 66, "Execute"

  return order, command
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

// When two or more consecutive named function parameters share a type, you can omit the type from all but the last.
func add(x, y int) int {
	return x + y
}

// A function can return any number of results.
func results(x, y string) (string, string) {
	return y, x
}

// A return statement without arguments returns the named return values. This is known as a "naked" return.
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
