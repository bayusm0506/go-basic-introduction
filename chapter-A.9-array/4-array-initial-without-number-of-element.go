package main

import "fmt"

func main() {
	var number = [...]int{2, 3, 4, 4, 5}

	fmt.Println("Array data \t:", number)
	fmt.Println("Total element \t:", len(number))
}
