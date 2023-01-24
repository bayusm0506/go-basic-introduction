package main

import "fmt"

type student struct {
	name string
	grad int
}

func (s student) changeName1(name string) {
	fmt.Println("----> on changeName1, name changed to : ", name)
}

func (s *student) changeName2(name string) {
	fmt.Println("----> on changeName2, name changed to : ", name)
	s.name = name
}

func main() {
	var s1 = student{"Jhon Doe", 22}
	fmt.Println("s1 before", s1.name)

	s1.changeName1("Rambo")
	fmt.Println("s1 after changeName1", s1.name)

	s1.changeName2("Elon")
	fmt.Println("s1 after changeName2", s1.name)
}
