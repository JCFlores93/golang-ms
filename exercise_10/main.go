package main

import (
	"fmt"

	"github.com/JCFlores93/golang-ms/exercise_10/dog"
)

type canine struct {
	name string
	age  int
}

func main() {
	fido := canine{
		name: "Fido",
		age:  dog.Years(10),
	}
	fmt.Println(fido)
}
