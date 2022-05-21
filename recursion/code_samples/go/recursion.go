package main

import (
	"fmt"
)

func recursion(n int) int {
	if n == 1 {
		return n
	}
	fmt.Println(n)
	_ = recursion(n - 1)
	return 0
}

func main() {
	x := recursion(10)
	fmt.Println(x)
}
