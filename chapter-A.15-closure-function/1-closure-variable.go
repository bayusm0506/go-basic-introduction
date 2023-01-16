package main

import "fmt"

func main() {
	var getMinMax = func(n []int) (int, int) {
		var min, max int
		for i, e := range n {
			switch {
			case i == 0:
				max, min = e, e
			case i > max:
				max = e
			case i < min:
				min = e
			}
		}
		return min, max
	}

	var numbers = []int{2, 5, 6, 3, 7, 9}
	var min, max = getMinMax(numbers)
	fmt.Printf("Data : %v\n min : %v\n max : %v\n", numbers, min, max)
}
