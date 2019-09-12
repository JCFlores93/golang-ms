package main

import "fmt"

func main() {
	// answer 1
	// Write a program that prints a number in decimal, binary, and hex
	// x := 5
	// fmt.Printf("%d\t%b\t%x\t", x, x, x)

	// answer 2
	// a := (42 == 42)
	// b := (42 <= 43)
	// c := (42 >= 43)
	// d := (42 != 43)
	// e := (42 < 43)
	// f := (42 > 43)
	// fmt.Println(a, b, c, d, e, f)

	// answer 3
	// const (
	// 	a     = 42
	// 	b int = 43
	// )
	// fmt.Println(a, b)

	// answer 4
	// var x int = 5
	// fmt.Printf("%d\t%b\t%x\t\n", x, x, x)
	// x <<= 1
	// fmt.Printf("%d\t%b\t%x\t", x, x, x)

	// answer 5
	// a := `here is something
	// as
	// a
	// raw string
	// literal
	// "you see"
	// another thing`
	// fmt.Println(a)

	// answer 6
	const (
		a = 2017 + iota
		b = 2017 + iota
		c = 2017 + iota
		d = 2017 + iota
	)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
