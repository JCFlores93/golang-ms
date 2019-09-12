package main

import "fmt"

func main() {
	// arrays()
	// slices()
	maps()
}

func arrays() {
	var x [5]int
	fmt.Println(x)
	x[4] = 42
	fmt.Println(x)
}

func slices() {
	// x := type{values} //composite literal
	// a slice allows you to group the same type
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x)
	fmt.Println(len(x))
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Println(x[1:3])
	fmt.Println(x[2:])
	x = append(x, 77, 88, 99, 1014)
	fmt.Println(x)
	y := []int{234, 456, 678, 987}
	y = append(x, y...)
	fmt.Println(y)
	// deleting from array
	x = append(x[:2], x[4:]...)
	fmt.Println(x)

	z := make([]int, 10, 100)
	fmt.Println(z)
	fmt.Println(len(z))
	fmt.Println(cap(z))

	z1 := make([]int, 5, 9)
	fmt.Println(z1)
	fmt.Println(len(z1))
	fmt.Println(cap(z1))
	z1 = append(z1, 1, 2, 3, 4, 5)
	fmt.Println(z1)
	fmt.Println(len(z1))
	fmt.Println(cap(z1))
	z1 = append(z1, 1, 2, 3, 4, 5, 7, 8, 9, 2)
	fmt.Println(z1)
	fmt.Println(len(z1))
	fmt.Println(cap(z1))

	jb := []string{"Jane", "Bond", "Chocolate", "Martini"}
	fmt.Println(jb)
	mp := []string{"Miss", "Moneypenny", "Strawberry", "Hazelnut"}
	fmt.Println(mp)
	xp := [][]string{jb, mp}
	fmt.Println(xp)
}

func maps() {
	m := map[string]int{
		"James":   32,
		"MsMoney": 27,
	}
	fmt.Println(m)
	v, ok := m["James"]
	fmt.Println(v)
	fmt.Println(ok)
	// check that element
	if v, ok := m["James"]; ok {
		fmt.Println(v)
		fmt.Println(ok)
	}

	m["jean"] = 33

	for k, v := range m {
		fmt.Println(k)
		fmt.Println(v)
	}
	if v, ok := m["James"]; ok {
		fmt.Println("value:", v)
		delete(m, "James")
	}
}
