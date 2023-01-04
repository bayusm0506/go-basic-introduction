package main

import "fmt"

func main() {
	var fruits = [4]string{"banana", "apple", "melon", "pineaple"}

	for _, fruit := range fruits {
		fmt.Printf("elemen %s\n", fruit)
	}
}
