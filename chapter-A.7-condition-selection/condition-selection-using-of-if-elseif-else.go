package main

import "fmt"

func main() {
	var point = 8

	if point == 10 {
		fmt.Println("Perfect")
	} else if point > 5 {
		fmt.Println("Good")
	} else if point == 4 {
		fmt.Println("Bad")
	} else {
		fmt.Println("Danger")
	}
}
