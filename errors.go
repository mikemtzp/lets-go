package main

import (
	"errors"
	"fmt"
	"strconv"
)

func error() {
	// Atoi converts a stringified number to an interger
	_, err := strconv.Atoi("42b")
	if err != nil {
		fmt.Println("couldn't convert:", err)
		// because "42b" isn't a valid integer, we print:
		// couldn't convert: strconv.Atoi: parsing "42b": invalid syntax
		// Note:
		// 'parsing "42b": invalid syntax' is returned by the .Error() method
		return
	}
	// if we get here, then
	// i was converted successfully
}

// Because errors are just interfaces, you can build your own custom types that implement the error interface
type userError struct {
  name string
}

func (e userError) Error() string {
  return fmt.Sprintf("%v has a problem with their account", e.name)
}

	// The Go standard library provides an "errors" package that makes it easy to deal with errors.
	var myErr = errors.New("something went wrong")


