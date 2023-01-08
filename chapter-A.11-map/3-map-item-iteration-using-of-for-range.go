package main

import "fmt"

func main() {
	var duck = map[string]int{
		"january":  10,
		"february": 25,
		"march":    45,
		"april":    23,
	}

	for key, val := range duck {
		fmt.Println(key, " \t", val)
	}
}
