package main

import "fmt"

func main() {
	dst := make([]string, 3)
	src := []string{"melon", "banana", "orange", "apple"}
	n := copy(dst, src)

	fmt.Println(dst)
	fmt.Println(src)
	fmt.Println(n)

	dst = []string{"potatao", "banana", "melon"}
	src = []string{"orange", "pineaple"}
	n = copy(dst, src)

	fmt.Println(dst)
	fmt.Println(src)
	fmt.Println(n)
}
