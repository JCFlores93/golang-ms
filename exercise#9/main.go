package main

import (
	"fmt"
)

func main() {
	// exercise1()
	exercise2()
}

func exercise1() {
	c := make(chan int)
	go func() {
		c <- 42
	}()
	fmt.Println(<-c)

	// buffered channel
	d := make(chan int, 1)
	d <- 500
	fmt.Println(<-d)
}

func exercise2() {
	// send-only type
	// cs := make(chan<- int)
	// Bidirecctional channel
	cs := make(chan int)

	go func() {
		cs <- 42
	}()
	fmt.Println(<-cs)

	fmt.Printf("------\n")
	fmt.Printf("cs\t%T\n", cs)

	// receive-only type
	// cr := make(<-chan int)
	cr := make(chan int)

	go func() {
		cr <- 42
	}()
	fmt.Println(<-cr)

	fmt.Printf("------\n")
	fmt.Printf("cr\t%T\n", cr)
}

func exercise3() {
	c := gen()
	receive(c)
	fmt.Println("about to exit")
}

func gen() <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()
	return c
}
func receive(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
}
