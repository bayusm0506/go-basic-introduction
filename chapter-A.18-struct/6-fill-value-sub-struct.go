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
	var p1 = person{name: "Jhon", age: 20}
	var s1 = student{person: p1, grade: 3}

	fmt.Println("Name : ", s1.name)
	fmt.Println("Age : ", s1.age)
	fmt.Println("Grade : ", s1.grade)
}
