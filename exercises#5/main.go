package main

import "fmt"

func main() {
	//exercise1()
	exercise2()
}

type persona struct {
	first      string
	last       string
	favFlavors []string
}

type vehicle struct {
	doors string
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func exercise3() {
	s := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "James",
		friends: map[string]int{
			"Moneypenny": 555,
			"Q":          777,
			"M":          888,
		},
		favDrinks: []string{
			"Martini",
			"Water",
		},
	}
	fmt.Println(s.first)
	fmt.Println(s.friends)
	fmt.Println(s.favDrinks)

	for k, v := range s.friends {
		fmt.Println(k, v)
	}

	for i, val := range s.favDrinks {
		fmt.Println(i, val)
	}
}

func exercise2() {
	v1 := truck{
		vehicle: vehicle{
			doors: "1",
			color: "blue",
		},
		fourWheel: true,
	}

	v2 := sedan{
		vehicle: vehicle{
			doors: "4",
			color: "red",
		},
		luxury: true,
	}

	fmt.Println(v1)
	fmt.Println(v2)
	fmt.Println(v1.fourWheel)
	fmt.Println(v2.luxury)

}
func exercise1() {
	p1 := persona{
		first: "jean",
		last:  "flores",
		favFlavors: []string{
			"chocolate",
			"martini",
			"rum and coke",
		},
	}

	p2 := persona{
		first: "jean",
		last:  "flores",
		favFlavors: []string{
			"strawberry",
			"vanilla",
			"capuccino",
		},
	}

	x := []persona{
		p1, p2,
	}
	for i, v := range x {
		fmt.Printf("%T\t%T\t", v, i)
		// fmt.Println(i)
		fmt.Println(v)
	}

	x1 := map[string][]persona{
		"last_name": x,
	}
	for i, v := range x1 {
		fmt.Printf("%T\t%T\t", v, i)
		// fmt.Println(i)
		fmt.Println(v)
	}

}
