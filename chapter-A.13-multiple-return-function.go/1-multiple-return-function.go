package main

import (
	"fmt"
	"math"
)

func calculate(d float64) (float64, float64) {
	// calculate area
	var area = math.Pi * math.Pow(d/2, 2)

	// calculate arround
	var circumference = math.Pi * d

	// return 2 value
	return area, circumference
}

func main() {
	var diameter float64 = 15
	var area, circumference = calculate(diameter)

	fmt.Println("Circle area\t\t: %.2f \n", area)
	fmt.Println("Circumference\t: %.2f \n", circumference)
}
