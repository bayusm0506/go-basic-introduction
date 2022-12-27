package main

import "fmt"

func main() {
	name := new(string)

	fmt.Println(name)  // result hexadecimal
	fmt.Println(*name) // result empty string""
}
