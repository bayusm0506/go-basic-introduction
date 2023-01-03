package main

import "fmt"

func main() {
	var numbers1 = [2][3]int{[3]int{3, 2, 3}, [3]int{3, 4, 5}}
	var numbers2 = [2][3]int{{3, 2, 3}, {3, 4, 5}}

	fmt.Println("Numbers1 is", numbers1)
	fmt.Println("Numbers2 is", numbers2)
}
