package main

import (
	"fmt"
	"runtime"
)

var x1 bool

var x int
var y float64

func main() {
	fmt.Println(x1)
	x1 = true
	fmt.Println(x1)
	a := 7
	b := 42
	fmt.Println(a == b)
	fmt.Println(a != b)

	// NUMERIC
	//x = int(2.34534)
	y = 42.34534
	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
	var a1 int8 = -128
	fmt.Println(a1)
	fmt.Printf("%T\n", a1)
	a1 = 127
	fmt.Println(a1)
	fmt.Printf("%T\n", a1)

	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)

	// STRING TYPE
	s := "Hello, playground"
	fmt.Println(s)
	fmt.Printf("%T\n", s)
	s = "Hello, Hawaii"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U ", s[i])
	}

	fmt.Println("")

	for i, v := range s {
		fmt.Printf("at index position %d we have hex %#x\n", i, v)
	}
}
