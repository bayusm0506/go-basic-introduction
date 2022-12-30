package main

import "fmt"

func main() {
outerLoop:

	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			if i == 3 {
				break outerLoop
			}

			fmt.Print("Matrix [", i, "][", j, "]", "\n")
		}
	}
}
