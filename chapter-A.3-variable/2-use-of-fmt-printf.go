package main

import "fmt"

func main() {
	var firstName string = "Bayu"

	var lastName string
	lastName = "Setra"

	fmt.Printf("Hello Bayu Setra!\n")
	fmt.Printf("Hello %s %s!\n", firstName, lastName)
	fmt.Println("Hello", firstName, lastName+"!")
}
