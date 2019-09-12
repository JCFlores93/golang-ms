package main

import "fmt"

type person struct {
	first string
	last  string
}
type secretAgent struct {
	person
	ltk bool
}
type human interface {
	speak()
}

// attach this function to this type
func (s secretAgent) speak() {
	fmt.Println("I am ", s.first, s.last, " - the secrent agent speak")
}

func (p person) speak() {
	fmt.Println("I am ", p.first, p.last, " - the person speak")
}

func bar1(h human) {
	switch h.(type) {
	// aseguramos que es del tipo
	// h.(person).first
	case person:
		fmt.Println("I was passed into bar", h.(person).first)
	case secretAgent:
		fmt.Println("I was passed into bar", h.(secretAgent).first)
	}
	fmt.Println("I was passed into bar1 human", h)
}
func main() {
	// foo()
	// bar("Jean")
	// s1 := woo("Moneypenny")
	// fmt.Println(s1)
	// x, y := mouse("Moneypenny", "Moneypenny")
	// fmt.Println(x)
	// fmt.Println(y)
	// foo1("Flores")
	// --- defer
	// defer foo2()
	// bar2()
	sa1 := secretAgent{
		person: person{
			"Jean",
			"Flores",
		},
		ltk: true,
	}
	fmt.Println(sa1)
	sa1.speak()

	p1 := person{
		"Tony",
		"Flores",
	}
	fmt.Println(p1)

	bar1(sa1)
	bar1(p1)

	// anonymous function
	func() {
		fmt.Println("Anonymous func ran")
		// passing parameters and execute
	}()
	func(x int) {
		fmt.Println("Anonymous func ran", x)
		// passing parameters and execute
	}(42)

	// func expressions
	f := func() {
		fmt.Println("My first func expression")
	}
	f()
	g := func(x int) {
		fmt.Println("My first func expression")
	}
	g(12)
	s1 := foos()
	fmt.Println(s1)

	barsx := bars()
	fmt.Printf("%T", barsx)
	barsx()

	// callbacks
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s := sum(ii...)
	fmt.Println(s)

	s3 := even(sum, ii...)
	fmt.Println("even numbers", s3)

}

func sum(xi ...int) int {
	fmt.Println("%T\n", xi)
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func even(f func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 == 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}

func foos() string {
	s := "Hell world"
	return s
}

func bars() func() int {
	return func() int {
		return 451
	}
}

// func (r receiver) identifier(parameters) (return(s)) { ... }
// Everythin in Go is PASS BY VALUE
func foo() {
	fmt.Println("hello from foo")
}

func bar(s string) {
	fmt.Println("Hello", s)
}

func woo(s string) string {
	return fmt.Sprint("Hello from woo", s)
}

func mouse(fn string, ln string) (string, bool) {
	return fmt.Sprint("Hello from woo", ln), true
}

// Variadic parameter
func foo1(s string, x ...int) int {
	fmt.Println(x)
	return x[0]
}

func foo2() {
	fmt.Println("foo")
}
func bar2() {
	fmt.Println("bar")
}
