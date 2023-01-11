package main

import (
	"fmt"
	"math"
)

func calculate(d float64) (area float64, circumference float64) {
	// calculate area
	area = math.Pi * math.Pow(d/2, 2)

	// calculate arround
	circumference = math.Pi * d

	return
}

func main() {
	var diameter float64 = 15
	var area, circumference = calculate(diameter)

	fmt.Println("Circle area\t\t: %.2f \n", area)
	fmt.Println("Circumference\t: %.2f \n", circumference)
}
