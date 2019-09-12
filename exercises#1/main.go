package main

import "fmt"

var x int = 42
var y string = "James Bond"
var z bool = true

type myowntype int

func main() {
	// exercise1()
	// exercise2()
	// s := exercise3()
	// fmt.Println(s)
	exercise4()
}

func exercise1() {
	x := 42
	y := "James Bond"
	z := true
	fmt.Println(x, y, z)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}

func exercise2() {
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}

func exercise3() string {
	s := fmt.Sprintf("%v\t%v\t%v", x, y, z)
	return s
}

func exercise4() {
	var x myowntype
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)
}

func exercise5() {
	var x myowntype
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)
}
