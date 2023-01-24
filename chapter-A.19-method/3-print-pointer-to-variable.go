package main

import "fmt"

type student struct {
	name  string
	grade int
}

func (s *student) sayHello() {
	fmt.Println("Hello : ", s.name)
}

func main() {
	var s1 = student{"Jhon Rambo", 21}
	s1.sayHello()

	var s2 = &student{"Elon", 23}
	s2.sayHello()
}
