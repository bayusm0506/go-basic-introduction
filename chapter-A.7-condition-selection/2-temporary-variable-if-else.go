package main

import "fmt"

func main() {
	var point = 8840.0

	if percent := point / 100; percent >= 100 {
		fmt.Printf("%.1f%s Perfect\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s Good\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s Bad\n", percent, "%")
	}
}
