package main

import "fmt"

func main() {
	var fruits = []string{"apple", "banana", "grape", "melon"}
	var bFruits = fruits[0:2]

	fmt.Println(cap(bFruits)) // 4
	fmt.Println(len(bFruits)) // 2

	fmt.Println(fruits)  // [apple banana grape melon]
	fmt.Println(bFruits) // [apple banana]

	var cFruits = append(bFruits, "papaya")

	fmt.Println(fruits)  // [apple banana papaya melon]
	fmt.Println(bFruits) // [apple banana]
	fmt.Println(cFruits) // [apple banana papaya]
}
