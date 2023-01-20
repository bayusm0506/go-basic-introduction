package main

import "fmt"

type student struct {
	name  string
	grade int
}

func main() {
	var s1 = student{}
	s1.name = "Jhon"
	s1.grade = 3

	var s2 = student{"Doe", 3}
	var s3 = student{s1.name, 2}
	var s4 = student{name: "Rambo"}

	fmt.Println("Student 1 : ", s1.name)
	fmt.Println("Student 2 : ", s2.name)
	fmt.Println("Student 3 : ", s3.name)
	fmt.Println("Student 4 : ", s4.name)
}
