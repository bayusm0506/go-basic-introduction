package main

import "fmt"

func main() {
	var duck map[string]int
	duck = map[string]int{}

	duck["january"] = 20
	duck["february"] = 70

	fmt.Println("January", duck["january"])
	fmt.Println("February", duck["february"])
}
