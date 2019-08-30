package main

import (
	"fmt"
)

// DECLARE teh variable "y"
// ASSIGN the value 43
// declare
var y = 43

// DECLARE there is a VARIABLESwith the identifier z
// and that the variable with the identifier z is of type int
// assigns the zerovalue oftype int to z
// false for booleans, "" for strings
// and nil for pointers, functions, interfaces,slices,channels and maps
var z int

func main() {
	x := 42
	fmt.Println(x)
	x = 99
	fmt.Println(x)
	y := 100 + 70
	fmt.Println(y)
}
