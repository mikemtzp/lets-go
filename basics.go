package main

import "fmt"

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

// constants can be computed as long as the computation can happen at compile time.
const firstName = "Lane"
const lastName = "Wagner"
const fullName = firstName + " " + lastName

// you cannot declare a constant that can only be computed at run-time like you can in JavaScript. This breaks:
// the current time can only be known when the program is running
// const currentTime = time.Now()

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

// Variables in Go are passed by value
// "Pass by value" means that when a variable is passed into a function,
// that function receives a copy of the variable.
func value(){
    x := 5
    increment(x)

    fmt.Println(x)
    // still prints 5,
    // because the increment function
    // received a copy of x
}

func increment(x int) { x++ }

// How to Ignore Return Values
// A function can return a value that the caller doesn't care about.
// We can explicitly ignore variables by using an underscore: _
func getPoint() (x int, y int) {
  return 3, 4
}

// ignore y value
func ignore() int {
  x, _ := getPoint()
  return x
}


// Conditionals
func ifElse() {
  height := 10

  // if statements in Go don't use parentheses around the condition:
  if height > 4 {
    fmt.Println("You are tall enough!")
  }

  if height > 6 {
    fmt.Println("You are super tall!")
  } else if height > 4 {
    fmt.Println("You are tall enough!")
  } else {
    fmt.Println("You are not tall enough!")
  }
}


func getLength(email string) int { return len(email) }
func initialStatement() {
  // An if conditional can have an "initial" statement.
  // The variable(s) created in the initial statement are only defined within the scope of the if body.
  // This is just some syntactic sugar that Go offers to shorten up code in some cases.
  // For example, instead of writing:
  email := "email@work.com"

  length := getLength(email)
  if length < 1 {
    fmt.Println("Email is invalid")
  }

  // Not only is this code a bit shorter, but it also removes length from the parent scope.
  // We can do:
  if length := getLength(email); length < 1 {
    fmt.Println("Email is invalid")
}
}

// Structs
// We use structs in Go to represent structured data.
// It's often convenient to group different types of variables together.
type bus struct {
  Make string
  Model string
  Height int
  Width int
}

// Nested Structs
// Structs can be nested to represent more complex entities:
type car struct {
  Make string
  Model string
  Height int
  Width int
  FrontWheel Wheel
  BackWheel Wheel
}

// Note 'Wheel' needs to be public so it can be available for 'car' struct
type Wheel struct {
  Radius int
  Material string
}

func strut() {
  myCar := car{}
  myCar.FrontWheel.Radius = 5
}

// Anonymous Structs
// An anonymous struct is just like a regular struct,
// but it is defined without a name and therefore cannot be referenced elsewhere in the code.
func anonymousStruct() struct{Make string; Model string} {
  myCar := struct {
    Make string
    Model string
  } {
    Make: "tesla",
    Model: "model 3",
  }

  return myCar
}

// Go doesn't support classes or inheritance in the complete sense.
// Embedded structs are just a way to elevate and share fields between struct definitions.
type automovile struct {
  make string
  model string
}

type truck struct {
  // "automovile" is embedded, so the definition of a
  // "truck" now also additionally contains all
  // of the fields of the car struct
  automovile
  bedSize int
}
