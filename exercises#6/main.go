package main

import (
	"fmt"
	"math"
)

func main() {
	// exercise1()
	// exercise2()
	// exercise3()
	// exercise4()
	// exercise5()
	// exercise6()
	// exercise7()
	exercise8()

}

func exercise1() {
	n := foo1()
	x, s := bar1()
	fmt.Println(n, x, s)
}
func foo1() int {
	return 42
}

func bar1() (int, string) {
	return 1984, "Big Brother is Watching"
}

func exercise2() {
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	n := foo2(ii...)
	fmt.Println(n)

	ii2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	n2 := bar2(ii2)
	fmt.Println(n2)
}

func foo2(nums ...int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	return total
}

func bar2(xi []int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func exercise3() {
	defer foo3()
	fmt.Println("Hello, playground")
}

func foo3() {
	defer func() {
		fmt.Println("foo DEFER ran")
	}()
	fmt.Println("foo ran")
}

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("I'm ", p.first, p.last)
}
func exercise4() {
	p := person{
		first: "Jean",
		last:  "Flores",
		age:   26,
	}
	p.speak()
}

type circle struct {
	radius float64
}

type square struct {
	length float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (s square) area() float64 {
	return s.length * s.length
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println("", s.area())
}
func exercise5() {
	circ := circle{
		radius: 12.345,
	}

	squa := square{
		length: 15,
	}

	info(circ)
	info(squa)
}

func exercise6() {
	func() {
		for i := 0; i < 100; i++ {
			fmt.Println(i)
		}
	}()

	fmt.Println("done")
}

var x int
var g func()

func exercise7() {
	f := func() {
		for i := 0; i < 3; i++ {
			fmt.Println(i)
		}
	}

	f()
	fmt.Printf("%T\n", f)

	fmt.Println(x)
	fmt.Printf("%T\n", x)

	g = f
	g()
	fmt.Printf("this is g %T\n", g)

	fmt.Println("done")
}

func exercise8() {
	f := sum()
	fmt.Println(f())
}

func sum() func() int {
	return func() int {
		return 0
	}
}

func exercise9() {
	g := func(xi []int) int {
		if len(xi) == 0 {
			return 0
		}
		if len(xi) == 1 {
			return xi[0]
		}
		return xi[0] + xi[len(xi)-1]
	}

	x := foo123(g, []int{1, 2, 3, 4, 5, 6})
	fmt.Println(x)
}

func foo123(f func(xi []int) int, ii []int) int {
	n := f(ii)
	n++
	return n
}

func exercise10() {
	f := foo1234()
	fmt.Println(f())
	fmt.Println(f())
	g := foo1234()
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
}

func foo1234() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
