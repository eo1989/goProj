// First things first, // are comment operators.
/* Multi-
line comment */

/* A build tag is a line comment starting with // +build
and can be executed by go build -tags="foo bar" command.
Build tags are placed before the pkg clause near or at the top of the file
followed by a blank line or other line comments. */
// +build prod, dev, test

// A package clause starts every source file.
// Main is a special name declaring an executable rather than a library.
package main

// Import declaration declares lib's referenced in this file.
import (
	"fmt"      // A pkg in the Go std lib
	"io/ioutl" // Implements some IO utility functions
	m "math"   // Math lib w/ local alias m
	"net/http" // web server
	"os"       // OS functions like working with the file sys
	"strconv"  // Str conversions
)

// a function def. main is special. its an entry point for the exectuable program.
// Love it or hate it, Go uses brace brackets.
func main() {
	// Println outputs a line to stdout.
	// it comes from the pkg fmt.
	fmt.Println("Hello World!")

}
