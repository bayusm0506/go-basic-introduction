package main

import "fmt"

func main() {
	// initial map horizontal
	var duck1 = map[string]int{"january": 20, "february": 70}

	var duck2 = map[string]int{
		"january":  20,
		"february": 70,
	}

	// or
	var duck3 = map[string]int{}
	var duck4 = make(map[string]int)
	var duck5 = *new(map[string]int)

	fmt.Println(duck1)
	fmt.Println(duck2)
	fmt.Println(duck3)
	fmt.Println(duck4)
	fmt.Println(duck5)
}
