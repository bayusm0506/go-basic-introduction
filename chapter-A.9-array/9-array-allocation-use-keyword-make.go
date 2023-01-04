package main

import "fmt"

func main() {
	var fruits = make([]string, 2)

	fruits[0] = "manggo"
	fruits[1] = "apple"

	fmt.Println(fruits)
}
