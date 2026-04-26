package main

import "fmt"

type Person struct {
	name string
	age  int
}

type People = Person

func main() {
	p1 := Person{"Jhon", 21}
	fmt.Println(p1)

	p2 := People{"Doe", 22}
	fmt.Println(p2)
}