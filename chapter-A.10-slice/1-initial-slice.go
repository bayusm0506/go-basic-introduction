package main

import "fmt"

func main() {
	var fruits = []string{"apple", "banana", "grape", "melon"}

	fmt.Println(fruits[0])

	var fruitsA = []string{"apple", "grape"} // slice
	var fruitsB = [2]string{"melon", "apple"}

	fmt.Println(fruitsA)
	fmt.Println(fruitsB)
}
