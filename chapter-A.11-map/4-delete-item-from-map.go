package main

import "fmt"

func main() {
	var duck = map[string]int{"january": 10, "february": 30}

	fmt.Println(len(duck))
	fmt.Println(duck)

	delete(duck, "january")

	fmt.Println(len(duck))
	fmt.Println(duck)
}
