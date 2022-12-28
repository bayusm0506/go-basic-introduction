package main

import "fmt"

func main() {
	const firstName string = "Jhon"
	fmt.Print("Hello ", firstName, "!\n")

	const lastName string = "Doe"
	fmt.Print("Nice to meet you ", lastName, "!\n")

	const (
		square         = "Box"
		isToday  bool  = true
		numeric  uint8 = 1
		floatNum       = 3.4
	)

	fmt.Println("============")
	fmt.Println(square)
	fmt.Println(isToday)
	fmt.Println(numeric)
	fmt.Println(floatNum)

	const (
		today string = "Tuesday"
		andNow
		isToday2 = true
	)

	fmt.Println(today)
	fmt.Println(andNow)
	fmt.Println(isToday2)

	fmt.Println("============")

	const one, two = 1, 2
	const t, f string = "Three", "Four"

	fmt.Println(one, two)
	fmt.Println(t, f)

}
