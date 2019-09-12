package main

import "fmt"

func main() {
	if true {
		fmt.Println("Hello, playground")
	}
	// right now x is scoped for the function
	if x := 42; x == 2 {
		fmt.Println("001")
	} else if x == 42 {
		fmt.Println("42a")
	} else {
		fmt.Println("our value was not 42")
	}
	// this wont work
	// fmt.Println(x)

	switch {
	case false:
		fmt.Println("this should not to print")
		// to continue even if the next case doesnt match
		//fallthrough
	}
}
