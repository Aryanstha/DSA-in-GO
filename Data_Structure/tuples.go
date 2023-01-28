package main

import (
	"fmt"
)

func powerSeries(a, b int) (int, int, int) {
	return a * a, b * b * b, a * b
}

func main() {

	var a, b, c int
	a, b, c = powerSeries(2, 3)
	fmt.Println(a, b, c)
}
