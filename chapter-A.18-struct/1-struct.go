package main

import "fmt"

type student struct {
	name  string
	grade int
}

func main() {
	var s1 student
	s1.name = "Jhon Doe"
	s1.grade = 3

	fmt.Println("Name : ", s1.name)
	fmt.Println("Grade : ", s1.grade)
}
