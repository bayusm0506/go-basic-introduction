package main

import "fmt"

func main() {
	var fruits = []string{"banana", "apple", "melon", "grape"}
	fmt.Println(len(fruits))
	fmt.Println(cap(fruits))

	var aFruits = fruits[0:3]
	fmt.Println(len(aFruits))
	fmt.Println(cap(aFruits))

	var bFruits = fruits[1:4]
	fmt.Println(len(bFruits))
	fmt.Println(cap(bFruits))
}
