package main

import (
	"fmt"
)

func main() {
	// exercise1()
	// exercise2()
	// exercise3()
	// exercise4()
	// exercise5()
	// exercise6()
	exercise7()
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

func exercise4() {
	q := make(chan int)
	c := gen1(q)

	receive2(c, q)

	fmt.Println("about to exit")
}
func receive2(c, q <-chan int) {
	for {
		select {
		case v := <-c:
			fmt.Println(v)
		case <-q:
			return
		}
	}
}
func gen1(q chan<- int) <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		q <- 1
		close(c)
	}()

	return c
}

func exercise5() {
	c := make(chan int)

	go func() {
		c <- 42
	}()
	v, ok := <-c
	fmt.Println(v, ok)

	close(c)

	v, ok = <-c
	fmt.Println(v, ok)
}

func exercise6() {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()
	for v := range c {
		fmt.Println(v)
	}
	fmt.Println("about to exit")
}

func exercise7() {
	c := make(chan int)

	for index := 0; index < 10; index++ {
		go func() {
			for i := 0; i < 10; i++ {
				c <- i
			}

		}()
	}

	for index := 0; index < 100; index++ {
		fmt.Println(index, <-c)
	}

	fmt.Println("about to exit")
}
