package main

import "fmt"

func main() {
	var left = false
	var right = true

	var leftAndRight = left && right
	fmt.Printf("Left && Right \t(%t) \n", leftAndRight)

	var leftOrRight = left || right
	fmt.Printf("Left || Right \t(%t) \n", leftOrRight)

	var leftReverse = !left
	fmt.Printf("!Left \t\t(%t) \n", leftReverse)
}
