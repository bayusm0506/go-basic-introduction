package main

import "fmt"

func main() {
	var point = 6

	switch point {
	case 8:
		fmt.Println("Perfect")
	case 7:
		fmt.Println("Good")
	default:
		fmt.Println("Not Bad")
	}
}
