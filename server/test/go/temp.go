package main

import (
	"fmt"
)

func main() {
	var (
		a int
		b int
	)
	fmt.Scanln(&a, &b)
	fmt.Println(Code(a, b))
}

func Code(a int, b int) int {
	return a + b
}
