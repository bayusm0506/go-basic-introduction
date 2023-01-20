package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	var s1 = struct {
		person
		grade int
	}{}

	s1.person = person{"Jhon", 21}
	s1.grade = 3

	fmt.Println("Name : ", s1.person.name)
	fmt.Println("Age : ", s1.person.age)
	fmt.Println("Grade : ", s1.grade)

	var s2 = struct {
		person
		grade int
	}{
		person: person{"Doe", 22},
		grade:  2,
	}

	fmt.Println("Name : ", s2.person.name)
	fmt.Println("Age : ", s2.person.age)
	fmt.Println("Grade : ", s2.grade)
}
