package main

import "fmt"

func main() {
	divideNumber(10, 2)
	divideNumber(4, 0)
	divideNumber(8, -4)
}

func divideNumber(m, n int) {
	if n == 0 {
		fmt.Println("Invalid divider, %d cannot divided by %d\n", m, n)
		return
	}

	var res = m / n
	fmt.Println("%d / %d = %d\n", m, n, res)
}
