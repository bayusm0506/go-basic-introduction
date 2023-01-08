package main

import "fmt"

func main() {
	var duck = map[string]int{"january": 10, "february": 30}
	var value, isExist = duck["march"]

	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("Item is not exist")
	}
}
