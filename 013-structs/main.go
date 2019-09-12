package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person
	ltk bool
}

func main() {
	p1 := person{
		first: "james",
		last:  "Bond",
	}
	p2 := person{
		first: "Miss",
		last:  "Moneypenny",
	}
	fmt.Println(p1.first)
	fmt.Println(p2.first)

	sa1 := secretAgent{
		person: person{
			first: "Miss",
			last:  "Moneypenny",
			age:   32,
		},
		ltk: true,
	}
	fmt.Println(sa1.person.age, sa1.person.first, sa1.person.age)
	p4 := struct {
		first string
		last  string
		age   int
	}{
		first: "Miss",
		last:  "Moneypenny",
		age:   32,
	}
	fmt.Println(p4)
}
