package main

import "fmt"

func findMax(numbers []int, max int) (int, func() []int) {
	var res []int

	for _, e := range numbers {
		if e <= max {
			res = append(res, e)
		}
	}

	return len(res), func() []int {
		return res
	}
}

func main() {
	var max = 3
	var numbers = []int{2, 4, 3, 5, 6, 1, 5, 7, 3, 7, 9, 3, 4, 8}
	var howMany, getNumbers = findMax(numbers, max)
	var theNumbers = getNumbers()

	fmt.Println("Numbers is \t:", numbers)
	fmt.Printf("Find \t: %d\n\n", max)

	fmt.Println("Found \t:", howMany)
	fmt.Printf("Value \t: %d", theNumbers)
}
