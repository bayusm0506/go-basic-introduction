package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	var student struct {
		person
		grade int
	}

	student.person = person{"Jhon", 20}
	student.grade = 2

	fmt.Println("Name : ", student.name)
	fmt.Println("Age : ", student.age)
	fmt.Println("Grade : ", student.grade)
}
