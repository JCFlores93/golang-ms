package main

import "fmt"

func main() {
	// for init; condition ; post{}
	for i := 0; i <= 100; i++ {
		fmt.Println("Hello, playground")
	}

	// nested looop
	for i := 0; i <= 10; i++ {
		for j := 0; j <= 3; j++ {
			fmt.Printf("The outer loop: %d\t the inner loop: %d\n", i, j)
		}
	}

	x := 1
	for x < 10 {
		fmt.Println(x)
		x++
	}
	fmt.Println("done")
}
