package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	// c := make(chan int)

	// go func() {
	// 	c <- 42
	// }()
	// // channels are blocked
	// fmt.Println(<-c)

	// usingWithFuncs()
	// commaokidiom()
	// fanin()
	// fanout()
	contextTest()

}

func wontwork() {
	// c := make(chan int)
	// c <- 42
	// // channels are blocked
	// fmt.Println(<-c)

	// // -------------------------

	// c := make(chan int, 1)
	// c <- 42
	// c <- 43
	// fmt.Println(<-c)
}

func work() {
	// c := make(chan int)

	// go func() {
	// 	c <- 42
	// }()
	// // channels are blocked
	// fmt.Println(<-c)

	// c := make(chan int, 1)
	// c <- 42
	// fmt.Println(<-c)

	// c := make(chan int, 2)
	// c <- 42
	// c <- 43
	// fmt.Println(<-c)
	// fmt.Println(<-c)

	// Only send values int, do not pull off
	// Send on the channel types int
	// d := make(chan<- int, 2)
	// d <- 42
	// d <- 43
	// fmt.Println(<-d)
	// fmt.Println(<-d)

	// Receives only channel type int
	// e := make(<-chan int, 2)
	// e <- 42
	// e <- 43
	// fmt.Println(<-e)
	// fmt.Println(<-e)
}

func usingWithFuncs() {
	c := make(chan int)
	// send
	go foo(c)
	// receive
	bar(c)
	fmt.Println("about to exit")
}

// send
func foo(c chan<- int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}

// receive
func bar(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
}

func selectChannel() {
	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	go send(eve, odd, quit)
	receive(eve, odd, quit)
	fmt.Println("about to exit")
}

func receive(e, o, quit <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("From the enven channel: ", v)
		case v := <-o:
			fmt.Println("From the odd channel: ", v)
		case v := <-quit:
			fmt.Println("From the default channel: ", v)
			return
		}

	}
}
func send(e, o chan<- int, q chan<- int) {
	for index := 0; index < 100; index++ {
		if index%2 == 0 {
			e <- index
		} else {
			o <- index
		}
	}
	close(e)
	close(o)
	q <- 0
}

func commaokidiom() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)

	go send1(even, odd, quit)
	receive1(even, odd, quit)
	fmt.Println("about to exit")
}

// send channel
func send1(even, odd chan<- int, quit chan<- bool) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	close(quit)
}

// receive channel
func receive1(even, odd <-chan int, quit <-chan bool) {
	for {
		select {
		case v := <-even:
			fmt.Println("the value received from the even channel:", v)
		case v := <-odd:
			fmt.Println("the value received from the odd channel:", v)
		case i, ok := <-quit:
			if !ok {
				fmt.Println("from comma ok bit", i)
				return
			} else {
				fmt.Println("from comma ok bit", i)
			}
		}
	}
}

func fanin() {
	even := make(chan int)
	odd := make(chan int)
	fanin := make(chan int)

	go send2(even, odd)

	go receive2(even, odd, fanin)

	for v := range fanin {
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}

// send channel
func send2(even, odd chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	close(even)
	close(odd)
}

// receive channel
func receive2(even, odd <-chan int, fanin chan<- int) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for v := range even {
			fanin <- v
		}
		wg.Done()
	}()

	go func() {
		for v := range odd {
			fanin <- v
		}
		wg.Done()
	}()

	wg.Wait()
	close(fanin)
}

func fanout() {
	c1 := make(chan int)
	c2 := make(chan int)

	go populate(c1)

	go fanOutIn(c1, c2)

	for v := range c2 {
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}

func populate(c chan int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}

func fanOutIn(c1, c2 chan int) {
	var wg sync.WaitGroup
	for v := range c1 {
		fmt.Println("value of v :", v)
		wg.Add(1)
		go func(v2 int) {
			fmt.Println("value of v inside:", v2)
			c2 <- timeConsumingWork(v2)
			wg.Done()
		}(v)
	}
	wg.Wait()
	close(c2)
}
func timeConsumingWork(n int) int {
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))
	return n + rand.Intn(1000)
}
func contextTest() {
	ctx := context.Background()

	fmt.Println("context:\t", ctx)
	fmt.Println("context err:\t", ctx.Err())
	fmt.Printf("context type:\t%T\n", ctx)
	fmt.Println("----------")

	ctx, cancel := context.WithCancel(ctx)

	fmt.Println("context:\t", ctx)
	fmt.Println("context err:\t", ctx.Err())
	fmt.Printf("context type:\t%T\n", ctx)
	fmt.Println("cancel:\t\t", cancel)
	fmt.Printf("cancel type:\t%T\n", cancel)
	fmt.Println("----------")

	cancel()

	fmt.Println("context:\t", ctx)
	fmt.Println("context err:\t", ctx.Err())
	fmt.Printf("context type:\t%T\n", ctx)
	fmt.Println("cancel:\t\t", cancel)
	fmt.Printf("cancel type:\t%T\n", cancel)
	fmt.Println("----------")
}
