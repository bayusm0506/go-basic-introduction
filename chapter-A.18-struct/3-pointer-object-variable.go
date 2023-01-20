package main

import "fmt"

type student struct {
	name  string
	grade int
}

func main() {
	var s1 = student{name: "Jhon", grade: 3}

	var s2 *student = &s1

	fmt.Println("Student 1, name : ", s1.name)
	fmt.Println("Student 4, name : ", s2.name)

	s2.name = "Doe"
	fmt.Println("Student 1, name : ", s1.name)
	fmt.Println("Student 4, name : ", s2.name)
}
