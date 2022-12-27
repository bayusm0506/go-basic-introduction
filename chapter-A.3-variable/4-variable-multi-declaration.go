package main

import "fmt"

func main() {
	var first, second, third string
	first, second, third = "Pertama", "Kedua", "Ketiga"

	var fourth, fifth, six string = "Keempat", "Kelima", "Keenam"

	seventh, eight, ninth := "Ketujuh", "Kedelapan", "Kesembilan"

	one, isTuesday, fourPointZero, say := 1, true, 4.0, "Hello"

	fmt.Println(first, second, third, fourth, fifth, six, seventh, eight, ninth)
	fmt.Println(one, isTuesday, fourPointZero, say)
}
