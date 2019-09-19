package main

import "fmt"

func main() {
	a := 42
	fmt.Println(a)
	fmt.Println(&a)

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)
	fmt.Printf("%T\n", *&a)
	b := &a
	*b = 43
	fmt.Println(a)
	fmt.Println(*b)

	mutableValue()
}

func mutableValue() {
	x := 0
	fmt.Println("x befor", &x)
	fmt.Println("x befor", x)
	foo(&x)
	fmt.Println("x after", &x)
	fmt.Println("x after", x)
}
func foo(y *int) {
	fmt.Println("y befor", y)
	fmt.Println("y befor", *y)
	*y = 43
	fmt.Println("y after", y)
	fmt.Println("y after", *y)
}
