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

	s1 := "Hello, playground"
	fmt.Println(s1)
	s1 = "Hello, Hawaii"
	fmt.Println(s1)
	fmt.Printf("%T\n", s1)

	bs1 := []byte(s1)
	fmt.Println(bs1)
	fmt.Printf("%T\n", bs1)

	for i := 0; i < len(s1); i++ {
		fmt.Printf("%#U ", s1[i])
	}

	fmt.Println("")

	for i, v := range s1 {
		fmt.Printf("at index position %d we have hex %#x\n", i, v)
	}

	// CONSTANTS
	const (
		d1 int     = 42
		e1 float64 = 42.78
		f1 string  = "James Bond"
	)
	fmt.Println(d1)
	fmt.Println(e1)
	fmt.Println(f1)
	fmt.Printf("%T\n", d1)
	fmt.Printf("%T\n", e1)
	fmt.Printf("%T\n", f1)

	// iota
	const (
		aaa = iota
		bbb
		c
	)
	const (
		d = iota
		e
		f
	)
	fmt.Println(aaa)
	fmt.Println(bbb)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)

	x := 4
	fmt.Printf("%d\t\t%b\n", x, x)
	y := x << 1
	fmt.Printf("%d\t\t%b", y, y)
	fmt.Printf("%d\t\t%b", 72, 72)

	kb := 1024
	mb := 1024 * kb
	gb := 1024 * mb

	fmt.Printf("%d\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t%b\n", gb, gb)

	const (
		_  = iota             // 0
		KB = 1 << (iota * 10) // 1 << (1 * 10)
		MB = 1 << (iota * 10) // 1 << (2 * 10)
		GB = 1 << (iota * 10) // 1 << (3 * 10)
		TB = 1 << (iota * 10) // 1 << (4 * 10)
	)

	fmt.Println("binary\t\t\t\tdecimal")
	fmt.Printf("%b\t\t\t", KB)
	fmt.Printf("%d\n", KB)
	fmt.Printf("%b\t\t", MB)
	fmt.Printf("%d\n", MB)
	fmt.Printf("%b\t", GB)
	fmt.Printf("%d\n", GB)
}
