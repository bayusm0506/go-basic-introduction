package main

import "fmt"

func main() {
	var fruits = [4]string{"banana", "apple", "manggo", "pineaple"}

	for i := 0; i < len(fruits); i++ {
		fmt.Printf("Element %d: %s\n", i, fruits[i])
	}
}
