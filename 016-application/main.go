package main

import (
	"encoding/json"
	"fmt"
	"sort"
)

type person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

type person1 struct {
	First string
	Last  string
	Age   int
}

func main() {

	p1 := person1{
		First: "James",
		Last:  "Bond",
		Age:   32,
	}

	p2 := person1{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
	}

	people := []person1{p1, p2}

	fmt.Println(people)

	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))

	s1 := `[{"First":"James","Last":"Bond","Age":32},{"First":"Miss","Last":"Moneypenny","Age":27}]`
	bs1 := []byte(s1)
	fmt.Printf("%T\n", s1)
	fmt.Printf("%T\n", bs1)

	var people1 []person

	err = json.Unmarshal(bs1, &people1)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("\nall of the data", people1)

	for i, v := range people1 {
		fmt.Println("\nPERSON NUMBER", i)
		fmt.Println(v.First, v.Last, v.Age)
	}
	sorted()
	sortedCustom()
	encryptPwd()
}

func sorted() {
	xi := []int{4, 7, 3, 42, 99, 18, 16, 56, 12}
	xs := []string{"James", "Q", "M", "Moneypenny", "Dr. No"}

	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println(xi)

	fmt.Println("------")
	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs)
}

type Person struct {
	First string
	Age   int
}

type ByName []Person

func (bn ByName) Len() int           { return len(bn) }
func (bn ByName) Swap(i, j int)      { bn[i], bn[j] = bn[j], bn[i] }
func (bn ByName) Less(i, j int) bool { return bn[i].First < bn[j].First }

func sortedCustom() {
	p1 := Person{"James", 32}
	p2 := Person{"Moneypenny", 27}
	p3 := Person{"Q", 64}
	p4 := Person{"M", 56}

	people := []Person{p1, p2, p3, p4}

	fmt.Println(people)
	sort.Sort(ByName(people))
	fmt.Println(people)

}

func encryptPwd() {
	s := `password123`
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)
	fmt.Println(bs)

	loginPword1 := `password1234`

	err = bcrypt.CompareHashAndPassword(bs, []byte(loginPword1))
	if err != nil {
		fmt.Println("YOU CAN'T LOGIN")
		return
	}

	fmt.Println("You're logged in")
}
