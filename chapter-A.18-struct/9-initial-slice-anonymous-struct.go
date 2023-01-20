package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	var allStudents = []struct {
		person
		grade int
	}{
		{person: person{"wick", 21}, grade: 2},
		{person: person{"ethan", 22}, grade: 3},
		{person: person{"bond", 23}, grade: 4},
	}

	for _, student := range allStudents {
		fmt.Println("Hello my name is", student.name, " Age is ", student.age)
	}
}
