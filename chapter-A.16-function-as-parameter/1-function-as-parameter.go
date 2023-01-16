package main

import (
	"fmt"
	"strings"
)

func filter(data []string, callback func(string) bool) []string {
	var result []string
	for _, each := range data {
		if filtered := callback(each); filtered {
			result = append(result, each)
		}
	}
	return result
}

func main() {
	var data = []string{"Jhon", "Rambo", "1990"}

	var dataContainSo = filter(data, func(each string) bool {
		return strings.Contains(each, "o")
	})

	var dataLength5 = filter(data, func(each string) bool {
		return len(each) == 5
	})

	fmt.Println("Real data \t\t:", data)

	fmt.Println("Filter any letters \"o\"\t:", dataContainSo)

	fmt.Println("Filter number of letters \"5\"\t:", dataLength5)
}
