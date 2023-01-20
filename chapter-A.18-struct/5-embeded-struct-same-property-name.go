package main

import "fmt"

type person struct {
	name string
	age  int
}

type student struct {
	person
	age   int
	grade int
}

func main() {
	var s1 = student{}
	s1.name = "Jhon"
	s1.age = 21
	s1.person.age = 20

	fmt.Println("Name : ", s1.name)
	fmt.Println("Age from student : ", s1.age)
	fmt.Println("Age from person : ", s1.person.age)
}
