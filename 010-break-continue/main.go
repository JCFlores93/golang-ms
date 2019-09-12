package main

import "fmt"

func main() {
	x := 0
	for {
		x++
		if x > 100 {
			break
		}
		if x%2 != 0 {
			continue
		}
		fmt.Println(x)
	}

	for i := 33; i <= 122; i++ {
		fmt.Printf("%v\t%x\t%#U\n", i, i, i)
	}

	// Printing ascii
	for i := 33; i <= 122; i++ {
		fmt.Printf("%v\t%#x\t%#U\n", i, i, i)
	}
}
