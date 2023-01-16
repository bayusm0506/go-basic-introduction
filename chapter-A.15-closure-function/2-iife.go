package main

import "fmt"

func main() {
	var numbers = []int{2, 4, 3, 5, 6, 1, 5, 7, 3, 7, 9, 3, 4, 8}

	var newNumbers = func(min int) []int {
		var r []int
		for _, e := range numbers {
			if e < min {
				continue
			}

			r = append(r, e)
		}

		return r
	}(3)

	fmt.Println("Original number :", numbers)
	fmt.Println("Filtered number :", newNumbers)
}
