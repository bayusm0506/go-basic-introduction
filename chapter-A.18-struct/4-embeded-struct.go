package main

import "fmt"

type person struct {
	name string
	age  int
}

type student struct {
	grade int
	person
}

func main() {
	var s1 = student{}
	s1.name = "Jhone"
	s1.age = 21
	s1.grade = 3

	fmt.Println("Name : ", s1.name)
	fmt.Println("Age : ", s1.age)
	fmt.Println("Grade : ", s1.grade)
}
