package main

import (
	"fmt"
	"time"
)

func main() {
	//exercise1()
	// exercise2()
	// exercise3()
	// exercise4()
	// exercise5()
	// exercise6()
	// exercise7()
	// exercise8()
	exercise9()
}

func exercise1() {
	for i := 1; i <= 10000; i++ {
		fmt.Printf("%v\t\n", i)
	}
}

func exercise2() {
	for i := 65; i <= 90; i++ {
		fmt.Println(i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t%#U\n", i)
		}
	}
}

func exercise3() {
	currentYear := time.Now().Year()
	yearOfBirth := 1993
	for yearOfBirth <= currentYear {
		fmt.Println(yearOfBirth)
		yearOfBirth++
	}
}

func exercise4() {
	currentYear := time.Now().Year()
	yearOfBirth := 1993
	for {
		if yearOfBirth > currentYear {
			break
		}
		fmt.Println(yearOfBirth)
		yearOfBirth++
	}
}

func exercise5() {
	for i := 10; i <= 100; i++ {
		fmt.Printf("When %v is divided by 4, the remainder (aka modulus) is %v\n", i, i%4)
	}
}

func exercise6() {
	x := "James Bond"

	if x == "James Bond" {
		fmt.Println(x)
	}
}

func exercise7() {
	x := "Moneypenny"

	if x == "Moneypenny" {
		fmt.Println(x)
	} else if x == "James Bond" {
		fmt.Println("BONDDONBONDONBOND", x)
	} else {
		fmt.Println("neither")
	}
}

func exercise8() {
	switch {
	case false:
		fmt.Println("should not print")
	case true:
		fmt.Println("should print")
	}
}

func exercise9() {
	favSport := "surfing"
	switch favSport {
	case "skiing":
		fmt.Println("go to the mountains!")
	case "swimming":
		fmt.Println("go to the pool!")
	case "surfing":
		fmt.Println("go to hawaii!")
	case "wingsuit flying":
		fmt.Println("what would you like me to say at your funeral?")
	}
}
