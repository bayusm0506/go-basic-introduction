package main

import (
	. "access-level/library"
	f "fmt"
)

func main() {
	var s1 = Student{"Jhon", 21}
	f.Println("Name : ", s1.Name)
	f.Println("Age : ", s1.Grade)
}
