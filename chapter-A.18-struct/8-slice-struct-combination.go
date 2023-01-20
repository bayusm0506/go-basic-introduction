package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	var allStudents = []person{
		{name: "Jhon", age: 20},
		{name: "Doe", age: 21},
		{name: "Wick", age: 22},
	}

	for _, student := range allStudents {
		fmt.Println("Hello my name is", student.name, "Age is ", student.age)
	}
}
